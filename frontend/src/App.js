import React, { useState } from 'react';
import './App.css';

function App() {
  const apiEndpoint = "http://localhost:9808/"
  // const apiEndpoint = "http://gourl.localhost/"
  const [long_url, setLongUrl] = useState('');
  const [custom_link, setCustomKey] = useState('');
  const [shortenedUrl, setShortenedUrl] = useState(''); // New state for shortened URL

  const handleSubmit = async (e) => {
    e.preventDefault();

    const requestBody = {
      long_url,
      custom_link,
    };

    try {
      const response = await fetch(apiEndpoint + 'create-short-url', {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json',
        },
        body: JSON.stringify(requestBody),
      });

      console.log(long_url)
      console.log(custom_link)
      console.log(JSON.stringify({ long_url, custom_link }))

      if (response.ok) {
        console.log(response)
        const data = await response.json();
        console.log(data)
        setShortenedUrl(data.short_url); // Update the state with the shortened URL
      } else {
        console.error('Failed to shorten URL');
      }
    } catch (error) {
      console.error('Error:', error);
    }
  };


  return (
    <div className="App">
      <header className="App-header">
        <h1>URL Shortener Service</h1>
        <p>Enter a URL and a custom key (max 8 characters) to shorten:</p>
        <form onSubmit={handleSubmit}>
          <input
            type="url"
            placeholder="Enter URL"
            value={long_url}
            onChange={(e) => setLongUrl(e.target.value)}
            required
          />
          <input
            type="text"
            placeholder="Custom Key (optional)"
            maxLength={8}
            value={custom_link}
            onChange={(e) => setCustomKey(e.target.value)}
          />
          <button type="submit">Shorten</button>
        </form>
        {shortenedUrl && (
          <div className="shortened-url">
            <p>Shortened URL: <a href={"localhost:9808/" + shortenedUrl} target="_blank" rel="noopener noreferrer">{"gourl.localhost/" + shortenedUrl}</a></p>
          </div>

        )}
      </header>
    </div>
  );
}

export default App;
