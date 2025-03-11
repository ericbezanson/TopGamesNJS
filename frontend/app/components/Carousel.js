// app/components/Carousel.js
'use client';

import React from 'react';
import Image from 'next/image';

const Carousel = ({ title, screenshots }) => {
  // Exclude the first screenshot for the carousel
  const carouselScreenshots = screenshots.slice(1);

  return (
    <div className="carousel-container mt-0">
      <h1 className="text-4xl font-bold mb-4">{title}</h1>
      {carouselScreenshots.length > 0 ? (
        <div className="flex space-x-8 space-y-0 overflow-x-auto scroll-smooth">
          {carouselScreenshots.map((screenshot, index) => (
            <div key={index} className="flex-shrink-0 w-[300px] h-[200px]">
              <a
                href={screenshot.url}
                target="_blank"
                rel="noopener noreferrer"
                className="block" // Make the <a> tag a block element
              >
                <Image
                  src={screenshot.url}
                  alt={`Screenshot ${index + 1}`}
                  width={300}
                  height={200}
                  className="rounded-lg shadow-lg object-cover cursor-pointer" // Add cursor-pointer
                />
              </a>
            </div>
          ))}
        </div>
      ) : (
        <p className="text-center text-gray-400">No additional screenshots available.</p>
      )}
    </div>
  );
};

export default Carousel;