"use client";

interface Props {
  heroes: any[];
}

export default function HeroList(props: Props) {
  const { heroes } = props;

  console.log({ heroes });
  if (!heroes) {
    return <div>Loading...</div>;
  }

  return (
    <>
      {heroes &&
        heroes.length > 0 &&
        heroes.map((hero: any) => {
          return (
            <div key={`${hero.hero_id}`}>
              <div key={`${hero.hero_name}`}>
                Hero Name: {hero.hero_name}
              </div>
            </div>
          );
        })}
    </>
  );
}
