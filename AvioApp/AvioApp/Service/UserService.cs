﻿using AvioApp.Model;
using AvioApp.Model.DTO;
using AvioApp.Repository;
using AvioApp.SupportClasses;
using AvioApp.SupportClasses.GEH.CustomExceptions;
using System.Text;

namespace AvioApp.Service
{
    public class UserService : IUserService
    {
        private readonly IUnitOfWork _unitOfWork;
        private readonly IJWTGenerator _jwtGenerator;

        public UserService(IUnitOfWork unitOfWork, IJWTGenerator jwtGenerator)
        {
            _unitOfWork = unitOfWork;
            _jwtGenerator = jwtGenerator;
        }

        public string Authenticate(CredentialsDTO credentials)
        {
            User? user = _unitOfWork.UserRepository.GetAll().FirstOrDefault(u => u.Email == credentials.Email);
            if (user == null)
            {
                throw new NotFoundException($"User with email {credentials.Email} does not exists!");
            }
            if (PasswordHasher.VerifyPassword(credentials.Password, user.Password, user.Salt))
            {
                return _jwtGenerator.GenerateToken(user, credentials.IsLoginPermanent);
            }
            throw new BadCredentialsException($"Incorrect password for user with email {credentials.Email}!");
        }

        public long Register(NewUserDTO user)
        {
            if (_unitOfWork.UserRepository.GetAll().Where(u => u.Email == user.Email).Any())
            {
                throw new DuplicateItemException($"Email {user.Email} is already taken!");
            }
            byte[] salt;
            var newUser = new User
            {
                Email = user.Email,
                Password = PasswordHasher.HashPassword(user.Password, out salt),
                Salt = salt,
                Role = "USER",
                FirstName = user.FirstName,
                LastName = user.LastName
            };
            newUser = _unitOfWork.UserRepository.Create(newUser);
            _unitOfWork.SaveChanges();
            return newUser.Id;
        }

        public void UpdateCode(string code, string email)
        {
            var user = _unitOfWork.UserRepository.GetAll().FirstOrDefault(u => u.Email == email);
            if (user == null)
            {
                throw new NotFoundException($"User with email: {email} does not exists!");
            }
            user.Code = Encoding.ASCII.GetBytes(code);
            _unitOfWork.SaveChanges();
        }
    }
}
