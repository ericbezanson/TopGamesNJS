// app/components/SearchBar.js
'use client';

import { useState } from 'react';

const SearchBar = ({ onSearch }) => {
  const [query, setQuery] = useState('');

  const handleInputChange = (e) => {
    setQuery(e.target.value);
  };

  const handleSubmit = (e) => {
    e.preventDefault();
    onSearch(query);
  };

  return (
    <form onSubmit={handleSubmit} className="mb-8">
      <input
        type="text"
        value={query}
        onChange={handleInputChange}
        placeholder="Search for a game..."
        className="p-2 rounded-lg w-full text-black"
      />
      <button type="submit" className="mt-2 p-2 bg-blue-500 rounded-lg">
        Search
      </button>
    </form>
  );
};

export default SearchBar;
