import Image from "next/image";

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
      <main className="max-w-6xl mx-auto">
        <h1 className="text-4xl font-bold text-center mb-10">
          IGDB MOST PLAYED GAMES
        </h1>

        {games.length > 0 ? (
          <div className="grid grid-cols-1 sm:grid-cols-2 md:grid-cols-3 gap-8">
            {games.map((game) => (
              <div
                key={game.id}
                className="bg-gray-800 p-4 rounded-lg shadow-lg flex flex-col items-center"
              >
                <Image
                  src={game.cover_url}
                  alt={game.name}
                  width={200}
                  height={250}
                  className="rounded-lg mb-4"
                />
                <h2 className="text-xl font-semibold">{game.name}</h2>
                <p className="text-sm text-gray-400 mt-2 text-center line-clamp-3">
                  {game.summary || "No summary available."}
                </p>
              </div>
            ))}
          </div>
        ) : (
          <p className="text-center text-gray-400">Loading...</p>
        )}
      </main>
    </div>
  );
}