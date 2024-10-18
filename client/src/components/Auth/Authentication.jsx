import React, { useState, useEffect } from 'react';
import { Navigate, useNavigate } from "react-router-dom"
import LoadingSpinner from '../Loading/LoadingSpinner';
import { useAuth } from '../../context/AuthContext';
import { useAuthActions } from '../../context/authActions';

const Authentication = ({currentPath }) => {
  const [isLogin, setIsLogin] = useState(true);
  const [email, setEmail] = useState('user03@gmail.com');
  const [username, setUsername] = useState('');
  const [password, setPassword] = useState('123456');
  const [confirmPassword, setConfirmPassword] = useState('');
  const [errorMessages, setErrorMessages] = useState([]);
  const [successMessage, setSuccessMessage] = useState('');
  const [loading, setLoading] = useState(false);
  const navigate = useNavigate();
  const { performLogin } = useAuthActions();

  useEffect(() => {
    if (currentPath === "/register") {
      setIsLogin(false);
    } else {
      setIsLogin(true);
    }
  }, [currentPath]);

  const toggleAuthMode = () => {
    setIsLogin(!isLogin);
    setErrorMessages([]);
    setSuccessMessage('');
  };

  const handleAuthentication = async (e) => {
    e.preventDefault();
    setErrorMessages([]);

    if (isLogin) {
      // Logic for login
      if (!email || !password) {
        setErrorMessages(['Please enter your email and password.']);
        return;
      }

      const isValidEmail = (email) => /^\S+@\S+\.\S+$/.test(email);
      if (!isValidEmail(email)) {
        setErrorMessages(['Please enter a valid email address.']);
        return;
      }
      setLoading(true);

      const data = {
        email,
        password
      }

      try {
        const response = await fetch("http://localhost:8080/api/v1/auth/login", {
          method: "POST",
          credentials: "include",
          headers: {
            "Content-Type": "application/json",
          },
          body: JSON.stringify(data),
        });
        
        const res = await response.json();
        console.log("Server response:", res);

        const result = res.data;
        
        setSuccessMessage("Login successful!");
        setEmail("");
        setPassword("");
        setErrorMessages([]);
        const timeout = setTimeout(() => {
          setLoading(false);
          performLogin(result);
          navigate("/");
        },2000);
  
      } catch (error) {
        setLoading(false);
        console.error("Error uploading data:", error);
      }
      
    } else {
      // Logic for signup
      if (!username) {
        setErrorMessages(['Please enter a username.']);
        return;
      }

      const isValidEmail = (email) => /^\S+@\S+\.\S+$/.test(email);
      if (!email || !isValidEmail(email) || !password || password !== confirmPassword) {
        setErrorMessages(['Please ensure all fields are filled and passwords match.']);
        return;
      }
      setLoading(true);

      const data = {
        username,
        email,
        password
      }

      try {
        const response = await fetch("http://localhost:8080/api/v1/auth/register", {
          method: "POST",
          headers: {
            "Content-Type": "application/json",
          },
          body: JSON.stringify(data),
        });
        
        if (!response.ok) {
          const errorMessage = await response.json();
          setErrorMessages(errorMessage.error);
          const timeout = setTimeout(() => {
            setLoading(false);
          },500);
          throw new Error(`Network response was not ok: ${errorMessage.error}`);
        }
  
        const result = await response.json();
        console.log("Server response:", result);
        setSuccessMessage("Click on the Login button to access your account.");
        setErrorMessages([]);
        setEmail("");
        setUsername("");
        setPassword("");
        setConfirmPassword("");
        const timeout = setTimeout(() => {
          setLoading(false);
          setIsLogin(true);
        },2000);
      } catch (error) {
        console.error("Error uploading data:", error);
      }
    }
  };

  return (
    <section id="hero" className="bg-color2 pt-28 flex items-center md:h-screen">
      <div className="container mx-auto flex flex-col-reverse md:flex-row items-center justify-center p-4 md:p-8 md:gap-12 overflow-hidden h-full max-w-screen-xl">
        <div className="w-full max-w-4xl bg-color4 rounded-lg shadow-lg overflow-hidden md:flex">
          {/* Left Section: Image for Login or Signup */}
          <div className={`w-full md:w-1/2 p-6 flex items-center justify-center ${isLogin ? 'order-1' : 'order-2'} bg-color1 transition-transform transform duration-300 ease-in-out`}>
            <img 
              src={`${isLogin ? '../src/assets/photo_upload.jpg' : '../src/assets/photo_upload_2.jpg'}`}
              alt="Auth Visual" 
              className="w-full h-auto object-cover rounded"
            />
          </div>

          {/* Right Section: Form */}
          { loading ? 
            <div className={`w-full md:w-1/2 p-4 md:p-8 flex flex-col justify-center ${isLogin ? 'order-2' : 'order-1'} bg-color4 transition-transform transform duration-300 ease-in-out`}>
              <LoadingSpinner />
            </div>
            :
            <div className={`w-full md:w-1/2 p-4 md:p-8 flex flex-col justify-center ${isLogin ? 'order-2' : 'order-1'} bg-color4 transition-transform transform duration-300 ease-in-out`}>
              <h2 className="text-xl md:text-2xl font-semibold text-color3 mb-6 text-center">
                {isLogin ? 'Login to Your Account' : 'Create Your Account'}
              </h2>

              {/* Display error message if any */}
              {errorMessages || successMessage ? (
                <div>
                  {errorMessages && errorMessages.length > 0 && (
                      <div className="mt-2 text-red-500 font-semibold text-sm text-center mb-4">
                          <ul>
                              {errorMessages.map((error, index) => (
                                  <li key={index}>{error}</li>
                              ))}
                          </ul>
                      </div>
                  )}
                  {successMessage && (
                    <div className="mt-2 text-green-500 font-semibold text-sm text-center mb-4">
                      {successMessage}
                    </div>
                  )}
                </div>
              ) : (
                ""
              )}

              <form className="space-y-6" onSubmit={handleAuthentication}>

                {/* User Input */}
                {!isLogin && (
                <div>
                  <label htmlFor="username" className="block text-sm font-medium text-color3">
                    User
                  </label>
                  <input
                    type="text"
                    id="username"
                    name="username"
                    value={username}
                    onChange={(e) => setUsername(e.target.value)}
                    required
                    className="w-full p-1.5 md:p-3 mt-2 border border-color1 rounded-lg focus:outline-none focus:ring-2 focus:ring-color1 focus:border-color1 transition-colors duration-200 ease-in-out"
                    placeholder="Enter your username"
                  />
                </div>
                )}

                {/* Email Input */}
                <div>
                  <label htmlFor="email" className="block text-sm font-medium text-color3">
                    Email Address
                  </label>
                  <input
                    type="email"
                    id="email"
                    name="email"
                    value={email}
                    onChange={(e) => setEmail(e.target.value)}
                    required
                    className="w-full p-1.5 md:p-3 mt-2 border border-color1 rounded-lg focus:outline-none focus:ring-2 focus:ring-color1 focus:border-color1 transition-colors duration-200 ease-in-out"
                    placeholder="Enter your email"
                  />
                </div>

                {/* Toggle between Login and Signup */}
                {isLogin && (
                  <div>
                    <label htmlFor="password" className="block text-sm font-medium text-color3">
                      Password
                    </label>
                    <input
                      type="password"
                      id="password"
                      name="password"
                      value={password}
                      onChange={(e) => setPassword(e.target.value)}
                      required
                      className="w-full p-1.5 md:p-3 mt-2 border border-color1 rounded-lg focus:outline-none focus:ring-2 focus:ring-color1 focus:border-color1 transition-colors duration-200 ease-in-out"
                      placeholder="Enter your password"
                    />
                  </div>
                )}
                {!isLogin && (
                  <div className='flex gap-3'>
                    <div>
                      <label htmlFor="password" className="block text-sm font-medium text-color3">
                        Password
                      </label>
                      <input
                        type="password"
                        id="password"
                        name="password"
                        value={password}
                        onChange={(e) => setPassword(e.target.value)}
                        required
                        className="w-full p-1.5 md:p-3 mt-2 border border-color1 rounded-lg focus:outline-none focus:ring-2 focus:ring-color1 focus:border-color1 transition-colors duration-200 ease-in-out"
                        placeholder="Enter your password"
                      />
                    </div>
                    <div>
                      <label htmlFor="confirm-password" className="block text-sm font-medium text-color3">
                        Confirm Password
                      </label>
                      <input
                        type="password"
                        id="confirm-password"
                        name="confirm-password"
                        value={confirmPassword}
                        onChange={(e) => setConfirmPassword(e.target.value)}
                        required
                        className="w-full p-1.5 md:p-3 mt-2 border border-color1 rounded-lg focus:outline-none focus:ring-2 focus:ring-color1 focus:border-color1 transition-colors duration-200 ease-in-out"
                        placeholder="Confirm your password"
                      />
                    </div>
                  </div>
                )}

                {/* Submit Button */}
                <div>
                  <button
                    type="submit"
                    className="w-full py-3 bg-color1 text-color4 rounded-lg font-semibold hover:bg-color3 transition-colors duration-300 ease-in-out"
                  >
                    {isLogin ? 'Login' : 'Sign Up'}
                  </button>
                </div>
              </form>

              {/* Toggle Between Login and Signup */}
              <div className="text-center mt-4">
                <span className="text-sm text-color3">
                  {isLogin ? "Don't have an account?" : "Already have an account?"}
                </span>
                <button 
                  onClick={toggleAuthMode} 
                  className="text-sm font-semibold text-color1 hover:underline ml-2"
                >
                  {isLogin ? 'Sign Up' : 'Login'}
                </button>
              </div>
            </div>
          }
        </div>
      </div>
    </section>
  );
};

export default Authentication;
