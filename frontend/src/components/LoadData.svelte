<script lang="ts">
  const getData = async (): Promise<string> => {
    const res = await fetch('/api', {
      headers: {'Access-Control-Allow-Origin': '*'}
    });
    return await res.json();
  };
</script>

{#await getData()}
  <p>Loading data...</p>
{:then people}
  {#each people as person}
    <p>Hello {person.name}!</p>
  {/each}
{:catch error}
  <p>Cannot load data</p>
  <p>{error.message}</p>
{/await}
