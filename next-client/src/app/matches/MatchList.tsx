'use client';

interface Props {
  matches: any[];
}

export default function MatchList(props: Props) {
  const { matches } = props;

  console.log({matches})
  if (!matches) {
    return <div>Loading...</div>;
  }

  return (
    <>
      {matches &&
        matches.length > 0 &&
        matches.map((match: any) => {
          return (
            <div key={`${match.match_id}`}>
              <div key={`${match.player1_id}`}>
                Player 1 ID: {match.player1_id}
              </div>
              <div key={`${match.player2_id}`}>
                Player 2 ID: {match.player2_id}
              </div>
              <div key={`${match.player1_hero}`}>
                Player 1 Hero: {match.player1_hero}
              </div>
              <div key={`${match.player2_hero}`}>
                Player 2 Hero: {match.player2_hero}
              </div>
              <div key={`${match.winner_id}`}>Winner ID: {match.winner_id}</div>
              <div key={`${match.loser_id}`}>Loser ID: {match.loser_id}</div>
              <div key={`${match.player1_decklist.String}`}>
                Player 1 Decklist: {match.player1_decklist.String}
              </div>
              <div key={`${match.player2_decklist.String}`}>
                Player 2 Decklist: {match.player2_decklist.String}
              </div>
              <div key={`${match.match_name}`}>
                Match Name: {match.match_name.String}
              </div>
              <div key={`${match.created_at}`}>
                Created At:{match.created_at}
              </div>
            </div>
          );
        })
      }
    </>
  )
}