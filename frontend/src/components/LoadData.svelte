<script lang="ts">
  interface People {
    id: number;
    name: string;
  }
  
  const getPeople = async (): Promise<People[]> => {
    const res = await fetch('/api', {
      headers: {'Access-Control-Allow-Origin': '*'}
    });
    return await res.json();
  };
</script>

{#await getPeople()}
  <p>Loading data...</p>
{:then people}
  {#each people as person}
    <p>Hello {person.name}!</p>
  {/each}
{:catch error}
  <p>Cannot load data</p>
  <p>{error.message}</p>
{/await}
