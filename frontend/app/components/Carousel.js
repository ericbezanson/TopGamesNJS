// app/components/Carousel.js
'use client';

import React from 'react';
import Image from 'next/image';

const Carousel = ({ title, screenshots }) => {
  // Exclude the first screenshot for the carousel
  const carouselScreenshots = screenshots.slice(1);

  return (
    <div className="carousel-container mt-8">
    <h1 className="text-4xl font-bold mb-4">{title}</h1>
      {carouselScreenshots.length > 0 ? (
        <div className="flex overflow-x-scroll space-x-4">
          {carouselScreenshots.map((screenshot, index) => (
            <div key={index} className="flex-shrink-0">
              <Image
                src={screenshot.url}
                alt={`Screenshot ${index + 1}`}
                width={300}
                height={200}
                className="rounded-lg shadow-lg"
              />
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
