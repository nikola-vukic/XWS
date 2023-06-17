import React from "react";
import classes from "./Login.module.css";
import { Link } from "react-router-dom";
import { useRef, useContext } from "react";
import AuthContext from "../store/auth-context";
import { useNavigate } from "react-router-dom";

const Login = () => {
  const emailRef = useRef();
  const permaRef = useRef();
  const pwRef = useRef();
  const authCtx = useContext(AuthContext);
  const navigate = useNavigate();
  const loginHandler = () => {
    event.preventDefault();

    fetch("http://localhost:8000/user/login", {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
      },
      body: JSON.stringify({
        username: emailRef.current.value,
        password: pwRef.current.value,
      }),
    })
      .then((res) => {
        if (res.ok) {
          return res.json();
        }
      })
      .then((data) => {
        console.log(data);
        const parsedJWT = parseJwt(data.accessToken);
        console.log(parsedJWT);
        authCtx.login(
          parsedJWT.userId,
          parsedJWT.role,
          parsedJWT.username,
          data.accessToken
        );
        setTimeout(() => {
          navigateLogin(parsedJWT.role);
        }, 100);
      })
      .catch((error) => {
        alert("Wrong credentials");
      });
  };

  function parseJwt(token) {
    if (!token) {
      return;
    }
    const base64Url = token.split(".")[1];
    const base64 = base64Url.replace("-", "+").replace("_", "/");
    return JSON.parse(window.atob(base64));
  }

  const navigateLogin = (role) => {
    if (role == "host") {
      navigate("/my-accommodations", { replace: true });
    } else {
      navigate("/accommodations", { replace: true });
    }
  };

  return (
    <div className={classes.centerDiv}>
      <div>
        <form className={classes.login}>
          <div className={classes.span}>
            <label>Username</label>
            <input className={classes.input} ref={emailRef}></input>
          </div>
          <div className={classes.span}>
            <label>Password</label>
            <input
              className={classes.input}
              ref={pwRef}
              type="password"
            ></input>
          </div>
          <div className={classes.keepLog}>
            <input type="checkbox" ref={permaRef}></input>
            <span>Keep me logged in </span>
          </div>
          <div className={classes.buttonContainerCenter}>
            <button className={classes.loginButton} onClick={loginHandler}>
              Login
            </button>
          </div>
          <span className={classes.registerSpan}>
            <Link to={"/register"}>Not a member? Register here!</Link>
          </span>
        </form>
      </div>
    </div>
  );
};

export default Login;
