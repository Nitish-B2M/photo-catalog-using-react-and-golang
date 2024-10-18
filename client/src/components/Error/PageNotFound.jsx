import React from 'react';
import { NavLink } from "react-router-dom"

const PageNotFound = () => {
    return (
        <div className="flex flex-col items-center justify-center h-screen bg-color2 text-color3 px-4">
            <h1 className="text-6xl md:text-8xl font-bold text-color5">404</h1>
            <p className="mt-4 text-lg md:text-xl text-center max-w-md">
                Oops! The page you are looking for does not exist.
            </p>
            <NavLink to="/" className="mt-6 text-base md:text-lg text-color6 underline hover:text-color1">
                Go Back Home
            </NavLink>
        </div>
    );
};

export default PageNotFound;
