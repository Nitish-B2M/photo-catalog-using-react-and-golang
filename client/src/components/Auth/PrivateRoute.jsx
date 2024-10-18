import React from 'react';
import { Navigate } from 'react-router-dom';
import { useAuth } from "../../context/AuthContext";

const PrivateRoute = ({ children }) => {
    const { state } = useAuth();
    console.log("private route: ",state);
    if (state && !state.isLoggedIn){
        return <Navigate to="/" replace />;
    }

    return children;
};

export default PrivateRoute;
