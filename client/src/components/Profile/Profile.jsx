import React, { useEffect, useState } from 'react';
import {unixToReadableTime} from '../../utils/utils';
import { formatDistanceToNow } from 'date-fns';

const Profile = () => {
    const [user, setUser] = useState({});
    const [loading, setLoading] = useState({});
    const [error, setError] = useState({});

    useEffect(() => {
        const fetchUserData = async () => {
          try {
            const response = await fetch("http://localhost:8080/api/v1/auth/user", {
              method: "GET",
              credentials: 'include',
              headers: {
                "Content-Type": "application/json",
              }
            });
            const data = await response.json();
            console.log("profile:", data);
            
            if (data.error && data.error.length > 0) {
              setError(data.error[0]);
              throw new Error(data.error[0]);
            }
            
            data.data.lastActive = formatDistanceToNow(new Date(data.data.lastActive * 1000), { addSuffix: true });
            data.data.createdAt = unixToReadableTime(data.data.created_at)
            data.data.updatedAt = unixToReadableTime(data.data.updated_at)
            setUser(data.data);
          } catch (err) {
            setError(err.message);
          } finally {
            setLoading(false);
          }
        };
    
        fetchUserData();
      }, []);

    const activities = [
      { id: 1, description: "Uploaded a new photo to the gallery", date: "October 3, 2024" },
      { id: 2, description: "Commented on a service review", date: "September 29, 2024" },
      { id: 3, description: "Updated profile picture", date: "September 15, 2024" },
    ];

    return (
        <div className="container mx-auto flex flex-col items-center justify-center p-8 md:gap-12 overflow-hidden h-full max-w-screen-xl">
            {/* Profile Information */}
            <h1 className="text-2xl font-bold">Profile Information</h1>
            {user ? (
              <div className="mt-4 bg-gray-100 p-6 rounded-lg shadow-lg w-full md:w-1/2  flex flex-col gap-2 border border-color1">
                <p className="font-bold">Username: <span className="font-normal">{user.username}</span></p>
                <p className="font-bold">Email: <span className="font-normal">{user.email}</span></p>
                <p className="font-bold">LastActive: <span className="font-normal">{user.lastActive}</span></p>
                <p className="font-bold">Joined: <span className="font-normal">{user.createdAt}</span></p>
              </div>
            ) : (
              error ? <p>{error}</p> : ""
            )}
            {/* Past Activity Section */}
            <h2 className="text-xl font-semibold mt-8">Past Activity</h2>
            <div className="mt-4 w-full md:w-2/3">
                {activities && activities.length > 0 ? (
                    <ul className="space-y-4">
                        {activities.map((activity) => (
                            <li key={activity.id} className="bg-white p-4 rounded-lg shadow-md">
                                <p>{activity.description}</p>
                                <p className="text-gray-500 text-sm">{activity.date}</p>
                            </li>
                        ))}
                    </ul>
                ) : (
                    <p className="text-gray-500">No recent activity.</p>
                )}
            </div>
        </div>
    );
};

export default Profile;
