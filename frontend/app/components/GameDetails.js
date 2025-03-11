// app/components/GameDetail.js
'use client';

import React from 'react';
import Image from 'next/image';

const GameDetails = ({ game }) => {
  // Function to format UNIX timestamp to 'Month Day, Year'
  const formatDate = (timestamp) => {
    const date = new Date(timestamp * 1000); // Convert to milliseconds
    return date.toLocaleDateString('en-US', {
      year: 'numeric',
      month: 'long',
      day: 'numeric',
    });
  };

  return (
    <div className="flex flex-col lg:flex-row items-center lg:items-start">
      {/* Cover Image */}
      <div className="w-full lg:w-1/5 flex justify-center lg:justify-start mb-6 lg:mb-0">
        <Image
          src={game.cover_url}
          alt={game.name}
          width={300}
          height={375}
          className="rounded-lg shadow-lg"
        />
      </div>

      {/* Game Details */}
      <div className="lg:ml-8 w-full lg:w-2/3 text-center lg:text-left">
        <h1 className="text-4xl font-bold mb-4">{game.name}</h1>
        <p className="text-gray-400 mb-2">
          Release Date: {formatDate(game.first_release_date)}
        </p>
        <p className="text-lg text-white">
          {game.summary || 'No summary available.'}
        </p>
      </div>
    </div>
  );
};

export default GameDetails;
