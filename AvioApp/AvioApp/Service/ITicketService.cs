﻿using AvioApp.Model.DTO;

namespace AvioApp.Service
{
    public interface ITicketService
    {
        void Buy(string flightId, int amount, string email);
        IEnumerable<TicketPreviewDTO> GetAll(string email);
    }
}
