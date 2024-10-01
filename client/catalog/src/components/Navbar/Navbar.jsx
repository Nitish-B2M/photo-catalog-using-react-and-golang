import React, { useState } from "react";
import { HiOutlineMenuAlt3 } from "react-icons/hi"

const Navbar = () => {

    const [isOpen, setIsOpen] = useState(false)
    const [activeSession, setActiveSession] = useState('home')
    const handleToggle = () => {
        setIsOpen(!isOpen)
    }

    const handleLogin = () => {
        console.log("::::::::::::::")
    }

    const navLinks = ( 
        <ul className="font-medium flex flex-col md:flex-row lg:space-x-8 sm:space-x-4 space-y-2 md:space-y-0 p-4 md:p-0">
            <li><a href="#home" className={`text-fontC ${activeSession === 'home' ? 'isActive' : ''}`}>Home</a></li>
            <li><a href="#about" className={`text-fontC ${activeSession === 'about' ? 'isActive' : ''}`}>About</a></li>
            <li><a href="#service" className={`text-fontC ${activeSession === 'service' ? 'isActive' : ''}`}>Service</a></li>
            <li><a href="#contact" className={`text-fontC ${activeSession === 'contact345re' ? 'isActive' : ''}`}>Contact</a></li>
        </ul>
    )

    return (
        <header className="bg-color5 text-white py-6 px-4 fixed top-0 left-0 right-0 z-10 overflow-x-hidden bg-blend-overlay">
            <div className="container max-w-screen-xl mx-auto flex justify-between items-center h-full">
                {/* logo */}
                <div className="text-white text-lg font-bold tracking-wider">
                    <a href="#">
                        {/* <img src="" alt="" /> */}
                        <h1 className="w-full font-optional size-6">Pixday</h1>
                    </a>
                </div>

                {/* navitems */}
                <div className="hidden md:flex flex-grow font-main justify-center">
                    <nav>{navLinks}</nav>
                </div>
                <div id="authentication" className="flex justify-between gap-6">
                    <button className="login font-semibold border-2 p-2 px-6 hover:bg-white hover:text-color5 transition-all ease-in-out duration-900 rounded-md"
                    onClick={handleLogin}
                    >Login</button>
                    <button className="register font-semibold border-2 p-2 px-6 hover:bg-white hover:text-color5 transition-all ease-in-out duration-900 rounded-md">Register</button>
                </div>

                <div className="block md:hidden">
                    <button 
                    onClick={handleToggle}
                    className={`text-white focus:outline-none ${isOpen ? 'border border-white' : ''}`}>
                        <HiOutlineMenuAlt3 className="size-6" />
                    </button>
                </div>
            </div>
        </header>
    );
};

export default Navbar;
