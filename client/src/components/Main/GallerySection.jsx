import React, { useEffect, useState } from 'react';
import "../../App.css";

const GallerySection = () => {
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState(null);
  const [galleryData, setGalleryData] = useState([]);

  useEffect(() => {

    const fetchImages = async () => {
      try {
        const response = await fetch("http://localhost:8080/api/v1/catalog/list", {
          method: "GET",
          headers: {
            "Content-Type": "application/json",
          },
          credentials: "include",
        });
        console.log("Response Status:", response.status);
        const data = await response.json();
        console.log("Fetched Data:", data);
        setGalleryData(data.data);
      } catch (err) {
        console.error("Error fetching images:", err);
        setError(err.message);
      } finally {
        setLoading(false);
      }
    };
  
    fetchImages();
  }, []);

  const baseUrl = "http://localhost:8080/";

  return (  
      <div className="container mx-auto flex flex-col items-center justify-center p-8 md:gap-8 overflow-hidden h-full max-w-screen-xl mt-20">
        <h1 className="text-2xl font-bold mb-2">Gallery</h1>
        <div className="column-1 gap-2 lg:gap-4 sm:columns-2 lg:columns-3 xl:columns-4 px-8">
  {galleryData && galleryData.length > 0 ? (
    galleryData.map((item) => (
      item.image ? ( 
        <div key={item.id} className="relative bg-white p-4 rounded-lg shadow-md h-fit mb-4 border border-color1 overflow-hidden group">
          <img
            src={`${baseUrl}${item.image.replace(/\\/g, '/')}`}
            alt={item.caption || "Gallery Image"}
            className="w-full h-auto"
            onError={(e) => (e.target.style.display = "none")} // Hide if image fails to load
          />
          <div className="absolute inset-0 bg-black bg-opacity-60 flex items-center justify-center text-white opacity-0 group-hover:opacity-100 transition-opacity duration-300">
            <div className="text-center">
              {item.caption && <p className="font-bold">{item.caption}</p>}
              {item.location && <p className="text-sm">{item.location}</p>}
            </div>
          </div>
        </div>
      ) : (
        <div key={item.id} className="w-full h-32 bg-gray-300 rounded-md flex items-center justify-center">
          <p className="text-gray-500">No image available</p>
        </div>
      )
    ))
  ) : (
    <p className="text-gray-500 bg-white p-4 rounded-lg shadow-md">No images available.</p>
  )}
</div>

    </div>
);

  
  


};

export default GallerySection;
