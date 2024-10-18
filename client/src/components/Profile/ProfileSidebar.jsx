import { FaUser, FaCog, FaSignOutAlt } from 'react-icons/fa';
import React, { useState } from "react";
import { AiFillCaretRight, AiFillCaretLeft} from "react-icons/ai";

const ProfileSidebar = ({ onLogout }) => {
  const [isProfileSidebarOpen, setIsProfileSidebarOpen] = useState(false);
    return (
        <div className={`fixed top-20 pt-7 text-color5 h-screen inset-0 bg-color2 z-10 transform ${isProfileSidebarOpen ? "translate-x-0" : "-translate-x-36"} transition-transform duration-300 ease-in-out lg:inset-auto lg:w-40 w-max md:p-3 p-5 flex flex-col items-center justify-start`}>
            <div className="text-md md:text-lg font-bold mb-5">Profile Menu</div>
            <ul className='overflow-hidden'>
                <li className="flex items-center mb-2 md:mb-4">
                    <FaUser className="mr-2" />
                    <a href="#profile" className="hover:underline" onClick={(isProfileSidebarOpen) => {setIsProfileSidebarOpen(false)}}>Profile</a>
                </li>
                <li className="flex items-center mb-2 md:mb-4">
                    <FaCog className="mr-2" />
                    <a href="#settings" className="hover:underline" onClick={(isProfileSidebarOpen) => {setIsProfileSidebarOpen(false)}}>Settings</a>
                </li>
                <li className="flex items-center mb-2 md:mb-4">
                    <FaSignOutAlt className="mr-2" />
                    <button onClick={onLogout} className="hover:underline">Logout</button>
                </li>
            </ul>
            <div className={`flex items-center fixed top-6 rounded-tr-lg rounded-br-lg -right-7 lg:-right-7 md:-right-8 z-10 bg-color2 p-2 px-1 font-main text-color5`}>
                <button onClick={() => setIsProfileSidebarOpen(!isProfileSidebarOpen)} >
                    {isProfileSidebarOpen ? <AiFillCaretRight className="text-xl" /> : <AiFillCaretLeft className="text-xl" />}
                </button>
            </div>
        </div>
    );
};


export default ProfileSidebar;
