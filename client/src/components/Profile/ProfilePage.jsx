import React, { useEffect, useState } from 'react';
import Sidebar from './ProfileSidebar';
import Profile from './Profile';
import Settings from './ProfileSetting';
import { useNavigate } from 'react-router-dom';
import { useAuth } from '../../context/AuthContext';
import { useAuthActions } from '../../context/authActions';

const ProfilePage = () => {
    const [activeSection, setActiveSection] = useState('profile');
    const { performLogout } = useAuthActions();
    const {user } = useAuth();

    useEffect(() => {
        const handleHashChange = () => {
            const hash = window.location.hash.replace('#', '');
            if (hash) {
                setActiveSection(hash);
            }
        };

        handleHashChange();
        window.addEventListener('hashchange', handleHashChange);
        return () => {
            window.removeEventListener('hashchange', handleHashChange);
        };
    }, []);

    const navigate = useNavigate();
    const handleLogout = async () => {
        try {
            if(user && !user.isLoggedIn){
                throw new Error("Illegal action")
            }
            const response = fetch("http://localhost:8080/api/v1/auth/logout",{
                credentials: "include"
            });
            performLogout();
            const data = await response.json();
            console.log("Fetched Data:", data);
            navigate("/");
        } catch (error) {
            console.log(error.Message)
        }
    };

    return (
        <div className="flex pt-20">
            <Sidebar onLogout={handleLogout} />
            <div className="flex-1">
                {activeSection === 'settings' ? <Settings /> : <Profile />}
            </div>
        </div>
    );
};

export default ProfilePage;
