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
      <table className="table-auto">
        <thead>
          <tr>
            <th>Player 1</th>
            <th>Player 1 Hero</th>
            <th>Player 1 Decklist</th>
            <th>Player 2</th>
            <th>Player 2 Hero</th>
            <th>Player 1 Decklist</th>
            <th>Created At</th>
          </tr>
        </thead>
        <tbody>
          {
            matches &&
            matches.length > 0 &&
            matches.map((match: any) => {
              return (
                <tr>
                  <td>match.player1_id</td>
                  <td>match.player1_hero</td>
                  <td>match.player1_decklist</td>
                  <td>match.player2_id</td>
                  <td>match.player1_hero</td>
                  <td>match.player1_decklist</td>
                  <td>match.created_at</td>
                </tr>
              );
            })
          }
        </tbody>
      </table>
    </>
  );
}