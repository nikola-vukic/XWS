import classes from "./Login.module.css";
import { Link, useNavigate } from "react-router-dom";
import React, { useState, useRef, useEffect } from "react";

const roles = ["User", "Host"];

const Register = () => {
  const nameRef = useRef();
  const lastNameRef = useRef();
  const emailRef = useRef();
  const cityRef = useRef();
  const countryRef = useRef();
  const addressRef = useRef();
  const usernameRef = useRef();
  const roleRef = useRef();
  const [pw, setPw] = useState();
  const [rePw, setRePw] = useState();
  const [validation, setValidation] = useState(true);
  const navigate = useNavigate();

  const changePwHandler = () => {
    setPw(event.target.value);
  };
  const changeRePwHandler = () => {
    setRePw(event.target.value);
  };

  useEffect(() => {
    if (rePw !== pw) {
      setValidation(false);
    } else {
      setValidation(true);
    }
  }, [rePw, pw]);

  const registerHandler = () => {
    event.preventDefault();

    fetch("https://localhost:5000/api/User/registration", {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
        Authorization: "Bearer ",
      },
      body: JSON.stringify({
        email: emailRef.current.value,
        password: pw,
        firstName: nameRef.current.value,
        lastName: lastNameRef.current.value,
      }),
    })
      .then((res) => {
        if (res.ok) {
          return res;
        } else if (res.status == 400) {
          throw new Error("Error");
        }
      })
      .then((data) => {
        if (data) console.log(data);
        navigate("/login");
      })
      .catch((error) => {
        alert(error);
      });
  };

  return (
    <div className={classes.centerDiv}>
      <div>
        <form className={classes.register}>
          <div className={classes.container}>
            <div>
              <div className={classes.span}>
                <label>First name</label>
                <input className={classes.input} ref={nameRef}></input>
              </div>
              <div className={classes.span}>
                <label>Last name</label>
                <input className={classes.input} ref={lastNameRef}></input>
              </div>
              <div className={classes.span}>
                <label>Country</label>
                <input className={classes.input} ref={countryRef}></input>
              </div>
              <div className={classes.span}>
                <label>City</label>
                <input className={classes.input} ref={cityRef}></input>
              </div>
              <div className={classes.span}>
                <label>Address</label>
                <input className={classes.input} ref={addressRef}></input>
              </div>
            </div>
            <div>
              <div className={classes.span}>
                <label>Username</label>
                <input className={classes.input} ref={usernameRef}></input>
              </div>
              <div className={classes.span}>
                <label>Email</label>
                <input className={classes.input} ref={emailRef}></input>
              </div>
              <div className={classes.span}>
                <label>Role: </label>
                <select className={classes.select} ref={roleRef}>
                  {roles.map((role, index) => {
                    return <option key={index}>{role}</option>;
                  })}
                </select>
              </div>
              <div className={classes.span}>
                <label>Password</label>
                <input
                  className={classes.input}
                  value={pw}
                  onChange={changePwHandler}
                ></input>
              </div>
              <div className={classes.span}>
                <label>Re enter password</label>
                <input
                  className={classes.input}
                  value={rePw}
                  onChange={changeRePwHandler}
                ></input>
              </div>
            </div>
          </div>
          <div className={classes.buttonContainerCenter}>
            {validation ? (
              <button className={classes.loginButton} onClick={registerHandler}>
                Register
              </button>
            ) : (
              <button
                className={classes.loginButtonDisabled}
                onClick={registerHandler}
                disabled
              >
                Register
              </button>
            )}
          </div>
          <div className={classes.buttonContainerCenter}>
            <span className={classes.registerSpan}>
              <Link to={"/login"}>Already a member? Sign in here!</Link>
            </span>
          </div>
        </form>
      </div>
    </div>
  );
};

export default Register;
