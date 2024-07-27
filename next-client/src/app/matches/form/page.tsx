async function getAllHeroes() {
  const data = await fetch("http://localhost:8000/api/heroes", {
    headers: {
      'Content-Type': 'application/json'
    }
  });

  if (!data.ok) {
    throw new Error("Failed to fetch matches");
  }

  const heroes = await data.json();

  console.log({ heroes });

  return heroes;
}

async function getGameFormats() {
  const data = await fetch("http://localhost:8000/api/formats", {
    headers: {
      'Content-Type': 'application/json'
    }
  });

    if (!data.ok) {
      throw new Error("Failed to fetch matches");
    }

    const formats = await data.json();

    console.log({ formats });

    return formats;
}

export default async function New() {
  const { heroes } = await getAllHeroes();
  const { formats } = await getGameFormats();

  return (
    <div>
      <div className="border-b border-gray-900/10 pb-12">
        <h2 className="text-base font-semibold leading-7 text-gray-900">
          Enter Matchmaking
        </h2>

        <div className="mt-10 grid grid-cols-1 gap-x-6 gap-y-8 sm:grid-cols-6">
          <div className="sm:col-span-3">
            <label
              htmlFor="format"
              className="block text-sm font-medium leading-6 text-gray-900"
            >
              Format
            </label>
            <div className="mt-2">
              <select
                id="format"
                name="format"
                autoComplete="format-name"
                className="block w-full rounded-md border-0 py-1.5 text-gray-900 shadow-sm ring-1 ring-inset ring-gray-300 focus:ring-2 focus:ring-inset focus:ring-indigo-600 sm:max-w-xs sm:text-sm sm:leading-6"
              >
                {formats &&
                  formats.length > 0 &&
                  formats.map((format) => {
                    return <option>{format.format_name}</option>;
                  })}
              </select>
            </div>
          </div>

          <div className="sm:col-span-3">
            <label
              htmlFor="hero"
              className="block text-sm font-medium leading-6 text-gray-900"
            >
              Hero
            </label>
            <div className="mt-2">
              <select
                id="hero"
                name="hero"
                autoComplete="hero-name"
                className="block w-full rounded-md border-0 py-1.5 text-gray-900 shadow-sm ring-1 ring-inset ring-gray-300 focus:ring-2 focus:ring-inset focus:ring-indigo-600 sm:max-w-xs sm:text-sm sm:leading-6"
              >
                {heroes &&
                  heroes.length > 0 &&
                  heroes.map((heroes) => {
                    return <option>{heroes.hero_name}</option>;
                  })}
              </select>
            </div>
          </div>

          <div className="sm:col-span-4">
            <label
              htmlFor="decklist"
              className="block text-sm font-medium leading-6 text-gray-900"
            >
              Decklist
            </label>
            <div className="mt-2">
              <input
                id="decklist"
                name="decklist"
                type="decklist"
                autoComplete="decklist"
                className="block w-full rounded-md border-0 py-1.5 text-gray-900 shadow-sm ring-1 ring-inset ring-gray-300 placeholder:text-gray-400 focus:ring-2 focus:ring-inset focus:ring-indigo-600 sm:text-sm sm:leading-6"
              />
            </div>
          </div>
        </div>
      </div>
    </div>
  );
}
