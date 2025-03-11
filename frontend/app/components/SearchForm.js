// app/search/SearchForm.js
'use client';

import { useState } from 'react';

export default function SearchForm({ onSearch }) {
  const [query, setQuery] = useState('');

  const handleSubmit = (e) => {
    e.preventDefault();
    onSearch(query);
  };

  return (
    <form onSubmit={handleSubmit} className="mb-8">
      <input
        type="text"
        value={query}
        onChange={(e) => setQuery(e.target.value)}
        placeholder="Search for a game..."
        className="p-2 rounded-l-lg"
      />
      <button type="submit" className="p-2 bg-blue-500 text-white rounded-r-lg">
        Search
      </button>
    </form>
  );
}