// app/components/Carousel.js
'use client';

import React from 'react';
import Image from 'next/image';
import Link from 'next/link';

const LinkCarousel = ({ title, games }) => {
  return (
    <div className="carousel-container mt-0 p-3">
      <h1 className="text-4xl font-bold mb-4">{title}</h1>
      {games.length > 0 ? (
        <div className="flex space-x-8 space-y-0 overflow-x-auto scroll-smooth">
          {games.map((game) => (
            <Link key={game.id} href={`/details/${game.id}`} className="flex-shrink-0 w-[200px]"> {/* Removed h-[350px] */}
              <div className="p-0 flex flex-col items-center cursor-pointer min-h-full flex-grow"> {/* Added min-h-full flex-grow */}
                <Image
                  src={game.cover_url}
                  alt={game.name}
                  width={200}
                  height={250}
                  className="rounded-lg mb-4 object-contain"
                />
                <h2 className="text-xl text-center font-semibold line-clamp-3 p-4">{game.name}</h2> {/* added line-clamp-3 */}
              </div>
            </Link>
          ))}
        </div>
      ) : (
        <p className="text-center text-gray-400">No additional screenshots available.</p>
      )}
    </div>
  );
};

export default LinkCarousel;