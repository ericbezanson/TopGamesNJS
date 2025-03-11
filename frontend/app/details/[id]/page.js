// app/gamedetail/[id]/page.js
import React from 'react';
import GameDetails from '@/app/components/GameDetails';
import Carousel from '@/app/components/Carousel';
import LinkCarousel from '@/app/components/LinkCarousel';
import Link from 'next/link';
import Image from "next/image";

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

  // Fetch details of similar games
  let validSimilarGames = [];

  try {
    const similarGames = await Promise.all(
      game.similar_games.map(async (similarGameId) => {
        try {
          const res = await fetch(`http://localhost:8080/gamedetail/${similarGameId}`);
          return await res.json();
        } catch (error) {
          console.error(`Error fetching details for game ID ${similarGameId}:`, error);
          return null;
        }
      })
    );

    // Filter out any failed fetches
    validSimilarGames = similarGames.filter((g) => g !== null);
  } catch(error){
      console.error("Error fetching similar games", error);
  }

  return (
    <div className="min-h-screen flex flex-col relative">
      {/* Background Section */}
      <div className="relative h-[500px]">
        <div
          className="absolute inset-0 bg-cover bg-center"
          style={{ backgroundImage: `url(${game.screenshots[0].url})` }}
        >
          <div className="absolute inset-0 bg-black opacity-60"></div>
        </div>
        <div className="relative z-10 p-8 sm:p-20 font-sans text-white">
          <GameDetails game={game} />
        </div>
      </div>

      {/* Carousel Section */}
      <div className="bg-black px-8 py-4">
        <Carousel title={"Screenshots"} screenshots={game.screenshots} />
      </div>
      {validSimilarGames.length > 0 && (
        <div className="bg-black px-8 py-4">
          <LinkCarousel title={"Similar Games"} games={validSimilarGames} />
        </div>
      )}
      <div className='bg-black p-8 h-10'>
        Footer
      </div>
      {/* Header Section */}
      <header className="absolute top-0 left-0 p-4 z-10">
        <Link href={"/"}>
          <Image
            src="/home.png"
            alt="Home"
            width={40}
            height={40}
            className="cursor-pointer"
          />
        </Link>
      </header>
    </div>
  );
}