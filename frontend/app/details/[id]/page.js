// app/gamedetail/[id]/page.js
import React from 'react';
import GameDetails from '@/app/components/GameDetails';
import Carousel from '@/app/components/Carousel';

export default async function GameDetailPage({ params }) {
  const { id } = params;

  let game = null;

  try {
    const res = await fetch(`http://localhost:8080/gamedetail/${id}`);
    game = await res.json();
  } catch (error) {
    console.error('Error fetching game details:', error);
  }

  if (!game) {
    return <div className="text-center text-white">Game not found</div>;
  }

  return (
    <div className="min-h-screen flex flex-col">
      {/* Background Section */}
      <div
        className="relative flex-grow"
        style={{
          backgroundImage: `url(${game.screenshots[0].url})`,
          backgroundSize: 'cover',
          backgroundPosition: 'center',
          opacity: 0.8,
        }}
      >
        <div className="absolute inset-0 bg-black opacity-60"></div>
        <div className="relative z-10 p-8 sm:p-20 font-sans text-white">
          <GameDetails game={game} />
        </div>
      </div>

      {/* Carousel Section */}
      <div className="bg-black p-8">
        <Carousel title={"Screenshots"} screenshots={game.screenshots} />
      </div>
    </div>
  );
}
