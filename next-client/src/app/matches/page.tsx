import MatchList from "./MatchList";

async function getMatches() {
  const data = await fetch("http://localhost:8000/api/matches", {
    headers: {
      'Content-Type': 'application/json',
    }
  });

  if (!data.ok) {
    throw new Error("Failed to fetch matches")
  }

  const matches = await data.json();

  console.log({ matches })

  return matches
}

export default async function Matches() {
  const data = await getMatches();

  return (
    <main>
      <h1>Matches</h1>
      <MatchList matches={data.matches} />
    </main>
  )

}