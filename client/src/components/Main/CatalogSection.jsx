import React, { useState, useEffect } from "react";
import { FiInfo } from "react-icons/fi";

const HeroSection = () => {
  const [image, setImage] = useState(null);
  const [imageSrc, setImageSrc] = useState(null);
  const [caption, setCaption] = useState("");
  const [location, setLocation] = useState("");
  const [tags, setTags] = useState([]);
  const [tagInput, setTagInput] = useState("");
  const [errorMessage, setErrorMessage] = useState("");
  const [successMessage, setSuccessMessage] = useState("");
  const [base64Image, setBase64Image] = useState("");

  const handleFileChange = async (e) => {
    const currFiles = e.target.files;
    if (currFiles.length > 0) {
      const file = currFiles[0];
      const src = URL.createObjectURL(file);
      setImageSrc(src);
      setImage(e.target.files[0]);
      
      const base64Image = await getBase64(file);
      setBase64Image(base64Image);
    }
  };

  const handleCaptionChange = (e) => {
    setCaption(e.target.value);
  };

  const handleLocationChange = (e) => {
    setLocation(e.target.value);
  };

  const getBase64 = (file) => {
    return new Promise((resolve, reject) => {
      const reader = new FileReader();
      reader.onloadend = () => {
        const base64String = reader.result;
        const base64Data = base64String.split(",")[1];
        resolve(base64Data);
      };
      reader.onerror = reject;
      reader.readAsDataURL(file);
    });
  };

  // Handle form submission
  const handleSubmit = async (e) => {
    e.preventDefault();

    const data = new FormData();
    data.append("caption", caption);
    data.append("location", location);
    data.append("tags", JSON.stringify(tags));
    if (image) {
      data.append('image', image); // Append the image file
    }
    // const data = {
    //   caption: caption,
    //   location: location,
    //   image: base64Image,
    //   tags: JSON.stringify(tags),
    //   publisher_id: "8afe596f-cdc0-4b4b-9e89-fe7c34303fed",
    // };

    console.log("Data being sent:", data);
    try {
      const response = await fetch("http://localhost:8080/api/v1/catalog/add", {
        method: "POST",
        credentials: "include",
        body: data,
      });

      if (!response.ok) {
        const errorMessage = await response.json();
        setErrorMessage(errorMessage.error);
        throw new Error(`Network response was not ok: ${errorMessage}`);
      }

      setSuccessMessage("Image uploaded successfully!");
      setErrorMessage("");
      const result = await response.json();
      console.log("Server response:", result);
      setImage(null);
      setCaption("");
      setLocation("");
      setTags([]);
      setImage("");
    } catch (error) {
      console.error("Error uploading data:", error);
    }
  };

  const handleTagInputChange = (e) => {
    setTagInput(e.target.value);
  };

  const handleKeyDown = (e) => {
    if (e.key === "Enter") {
      e.preventDefault();
      if (tagInput.trim()) {
        setTags([...tags, tagInput.trim()]);
        setTagInput("");
      }
    }
  };

  const removeTag = (tagToRemove) => {
    setTags(tags.filter((tag) => tag !== tagToRemove));
  };

  return (
    <section id="hero" className="bg-color2 pt-28 flex flex-col items-center md:h-screen">
      <h1 className="text-2xl font-bold mb-4">Upload Your Photo</h1>
      <div className="container mx-auto flex flex-col-reverse md:flex-row items-center justify-center p-8 md:gap-12 overflow-hidden h-full max-w-screen-xl">
        {/* Upload Form Section with Glassmorphism */}
        <div
          id="upload-form"
          className="flex flex-col items-center w-full md:w-3/5 bg-color4/40 backdrop-blur-lg md:p-6 rounded-lg shadow-md border border-color2"
        >
          <form
            onSubmit={handleSubmit}
            method="post"
            className="flex flex-col gap-4 w-full bg-color4/70 p-4 rounded-lg shadow-sm"
          >
            <input
              type="text"
              name="caption"
              value={caption}
              onChange={handleCaptionChange}
              placeholder="Caption"
              className="p-4 border border-color1 rounded bg-color4/80 text-color3 transition-colors duration-300 focus:outline-none focus:ring-2 focus:ring-color1 text-sm md:text-base" // Increased padding
            />
            <input
              type="text"
              name="location"
              value={location}
              onChange={handleLocationChange}
              placeholder="Location"
              className="p-4 border border-color1 rounded bg-color4/80 text-color3 transition-colors duration-300 focus:outline-none focus:ring-2 focus:ring-color1 text-sm md:text-base" // Increased padding
            />
            <div className="mb-4">
              <label
                htmlFor="file"
                className="block mb-2 text-sm font-medium text-color3 md:text-base"
              >
                Upload Image <span className="text-red-500">*</span>
              </label>
              <input
                type="file"
                name="image"
                onChange={handleFileChange}
                id="image"
                accept="image/*"
                className="block w-full text-sm text-gray-500 file:mr-4 file:py-2 file:px-4 file:rounded file:border-0 file:text-sm file:font-semibold file:bg-color1 file:text-color4 transition-all duration-300 hover:file:bg-color3 focus:outline-none focus:ring-2 focus:ring-color1"
                required
              />
            </div>

            <div className="mb-4">
              <label
                htmlFor="tags"
                className="block mb-2 text-sm font-medium text-color3 md:text-base"
              >
                Tags
              </label>
              <div className="flex items-center gap-2">
                {/* Tags display */}
                <div className="flex flex-wrap gap-2">
                  {tags.map((tag, index) => (
                    <span
                      key={index}
                      className="bg-color1 text-color4 px-2 py-1 rounded-lg flex items-center"
                    >
                      {tag}
                      <button
                        type="button"
                        className="ml-2 text-color4 hover:text-red-500"
                        onClick={() => removeTag(tag)}
                      >
                        &times;
                      </button>
                    </span>
                  ))}

                  {/* Input for adding tags */}
                  <input
                    type="text"
                    value={tagInput}
                    onChange={handleTagInputChange}
                    onKeyDown={handleKeyDown}
                    placeholder="Add a tag and press Enter"
                    className="p-2 border border-color1 rounded bg-color4/80 text-color3 transition-colors duration-300 focus:outline-none focus:ring-2 focus:ring-color1"
                  />
                </div>

                {/* Hint Icon with Tooltip */}
                <div className="relative group">
                  <FiInfo className="text-color3 cursor-pointer" size={18} />
                  {/* Tooltip */}
                  <div className="absolute left-1/4 transform -translate-x-3/4 mb-1 bottom-8 w-60 bg-color2 text-black text-sm px-3 py-2 rounded-md shadow-lg opacity-0 group-hover:opacity-100 transition-opacity duration-300">
                    You need to hit Enter to add a tag
                  </div>
                </div>
              </div>
            </div>

            {errorMessage || successMessage ? (
              <div>
                {errorMessage && (
                  <div className="mt-2 text-red-500 font-semibold">
                    {errorMessage}
                  </div>
                )}
                {successMessage && (
                  <div className="mt-2 text-green-500 font-semibold">
                    {successMessage}
                  </div>
                )}
              </div>
            ) : (
              ""
            )}

            <button
              type="submit"
              className="w-full py-3 bg-color1 text-color4 rounded-lg transition-all duration-300 hover:bg-color3 focus:outline-none focus:ring-2 focus:ring-color1 text-sm md:text-base" // Increased padding
            >
              Upload
            </button>
          </form>
        </div>

        {/* Upload Preview Section with Glassmorphism */}
        <div
          id="upload-preview"
          className="flex justify-center items-center w-full md:h-[440px] md:w-2/5 bg-color4/40 backdrop-blur-lg p-6 rounded-lg shadow-md border border-color2 mb-8 md:mb-0 flex-col"
        >
          {imageSrc ? (
            <>
              <img
                id="file-preview"
                src={imageSrc}
                alt="Preview"
                className="w-80 h-72 object-cover border border-color1 rounded-lg shadow-sm"
              />
              {caption && (
                <div className="mt-2 text-color3 font-semibold text-lg">
                  {caption}
                </div>
              )}
              {location && (
                <div className="mt-1 text-color3 text-sm">{location}</div>
              )}
            </>
          ) : (
            <div className="w-full md:h-full border border-color1 rounded-lg flex items-center justify-center text-color3 bg-color4/80 backdrop-blur-lg p-4">
              <span className="text-center text-sm md:text-lg">
                No Image Selected
              </span>
            </div>
          )}
        </div>
      </div>
    </section>
  );
};

export default HeroSection;
