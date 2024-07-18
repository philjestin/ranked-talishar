import HeroList from "./HeroList";

async function getHeroes() {
  const data = await fetch("http://localhost:8000/api/heroes", {
    headers: {
      "Content-Type": "application/json",
    },
  });

  if (!data.ok) {
    throw new Error("Failed to fetch matches");
  }

  const heroes = await data.json();

  console.log({ heroes });

  return heroes;
}

export default async function Heroes() {
  const data = await getHeroes();

  return (
    <main>
      <h1>Matches</h1>
      <HeroList heroes={data.heroes} />
    </main>
  );
}
