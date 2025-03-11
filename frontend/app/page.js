import Image from "next/image";
import Link from 'next/link';

export default async function Home() {
  let games = [];

  try {
    const res = await fetch("http://localhost:8080/games");
    games = await res.json();
  } catch (error) {
    console.error("Error fetching games:", error);
  }

  return (
    <div className="min-h-screen p-8 sm:p-20 font-sans bg-gray-900 text-white">
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
      <main className="max-w-6xl mx-auto">
        <h1 className="text-4xl font-bold text-center mb-10">
          IGDB MOST PLAYED GAMES
        </h1>

        {games.length > 0 ? (
          <div className="grid grid-cols-1 sm:grid-cols-2 md:grid-cols-3 gap-8">
            {games.map((game) => (
              <Link key={game.id} href={`/details/${game.id}`}>
                <div className="bg-gray-800 p-4 rounded-lg shadow-lg flex flex-col items-center cursor-pointer">
                  <Image
                    src={game.cover_url}
                    alt={game.name}
                    width={200}
                    height={250}
                    className="rounded-lg mb-4"
                  />
                  <h2 className="w-[300px] text-xl text-center font-semibold text-ellipsis overflow-hidden whitespace-nowrap">{game.name}</h2>
                </div>
              </Link>
            ))}
          </div>
        ) : (
          <p className="text-center text-gray-400">Loading...</p>
        )}
      </main>
    </div>
  );
}