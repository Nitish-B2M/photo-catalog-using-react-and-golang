import React, { useState } from "react";
import { HiOutlineMenuAlt3, HiX, HiUserAdd } from "react-icons/hi"
import { NavLink } from "react-router-dom"
import { FaUserCircle } from 'react-icons/fa';
import { useAuth } from "../../context/AuthContext";

const Navbar = () => {
    const { state } = useAuth();
    console.log(state,"::: navbarr")

    const [isOpen, setIsOpen] = useState(false)
    const handleToggle = () => {
        setIsOpen(!isOpen)
    }

    const navLinks = ( 
        <ul className="font-medium flex flex-col md:flex-row lg:space-x-8 sm:space-x-4 space-y-2 md:space-y-0 p-4 md:p-0">
            <li><NavLink to="/" className={({ isActive }) => isActive ? "text-fontC font-semibold isActive hover:text-color7" : "text-fontC hover:text-color7"} onClick={(isOpen) => {setIsOpen(false)}}>Home</NavLink></li>
            <li><NavLink to="/gallery" className={({ isActive }) => isActive ? "text-fontC font-semibold isActive" : "text-fontC hover:text-color7"} onClick={(isOpen) => {setIsOpen(false)}}>Gallery</NavLink></li>
            <li><NavLink to="/service" className={({ isActive }) => isActive ? "text-fontC font-semibold isActive" : "text-fontC hover:text-color7"} onClick={(isOpen) => {setIsOpen(false)}}>Service</NavLink></li>
            <li><NavLink to="/contact" className={({ isActive }) => isActive ? "text-fontC font-semibold isActive" : "text-fontC hover:text-color7"} onClick={(isOpen) => {setIsOpen(false)}}>Contact</NavLink></li>
        </ul>
    )

    return (
        <header className="bg-color1 text-color4 py-4 px-4 fixed top-0 h-20 left-0 right-0 z-10 overflow-x-hidden bg-blend-overlay border-b border-color3">
            <div className="container max-w-screen-xl mx-auto flex justify-between items-center h-full">
                <div className="text-color4 text-lg font-bold tracking-wider">
                    <NavLink to="/">
                        <h1 className="w-full font-optional size-6">Pixday</h1>
                    </NavLink>
                </div>

                <div className="hidden md:flex flex-grow font-main justify-center">
                    <nav>{navLinks}</nav>
                </div>
                <div id="authentication" className="hidden md:flex justify-between gap-6">
                    {state && state.isLoggedIn ? 
                        <NavLink to="/profile" className={({ isActive }) => isActive ? "font-semibold border-2 p-2 px-4 bg-color4 text-color1 transition-all ease-in-out duration-900 rounded-md flex flex-row gap-2 items-center" : "font-semibold border-2 p-2 px-4 hover:bg-color4 hover:text-color1 transition-all ease-in-out duration-900 rounded-md flex flex-row gap-2 items-center"}>
                            <FaUserCircle className="text-3xl" />
                            {state?.user?.username}
                        </NavLink>
                    :
                        <NavLink to="/auth" className={({ isActive }) => isActive ? "font-semibold border-2 p-2 px-4 bg-color4 text-color1 transition-all ease-in-out duration-900 rounded-md flex flex-row gap-2 items-center" : "font-semibold border-2 p-2 px-4 hover:bg-color4 hover:text-color1 transition-all ease-in-out duration-900 rounded-md flex flex-row gap-2 items-center"}>
                            <HiUserAdd />
                            Authenticate
                        </NavLink>
                    }
                </div>

                <div className="block md:hidden">
                    <button 
                    onClick={handleToggle}
                    className={`text-color4 focus:outline-none ${isOpen ? 'border border-color6' : ''}`}>
                        {isOpen ? <HiX className="size-6 " /> : <HiOutlineMenuAlt3 className="size-6" /> }
                    </button>
                </div>

                <div className={`md:hidden fixed right-0 top-20 text-color4 w-44 ${isOpen ? 'block' : 'hidden'} bg-color1`}>
                    <nav>{navLinks}</nav>
                    <div id="authentication" className="md:hidden justify-between gap-6">
                    {state && state.isLoggedIn ? 
                        <NavLink to="/profile" className="font-semibold p-2 px-4 border-t-2 flex flex-row gap-2 items-center" onClick={(isOpen) => {setIsOpen(false)}}>
                            <FaUserCircle className="text-xl" />{state?.user?.username}
                        </NavLink>
                    :
                        <NavLink to="/auth" className="font-semibold p-2 px-4 border-t-2 flex flex-row gap-2 items-center hover:text-color7" onClick={(isOpen) => {setIsOpen(false)}}>
                            <HiUserAdd className="text-xl" />
                        Authenticate
                        </NavLink>
                    }
                    </div>
                </div>
            </div>
        </header>
    );
};
export default Navbar;
