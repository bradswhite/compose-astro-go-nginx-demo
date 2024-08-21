<script lang="ts">
  const getData = async (): Promise<string> => {
    //const url = import.meta.env.PUBLIC_API_URL;
    const url = '/api';
    const res = await fetch(url, {
      headers: {'Access-Control-Allow-Origin': '*'}
    });
    return await res.json();
  };
</script>

{#await getData()}
  <p>Loading data...</p>
{:then data}
  <p>Hello {data.name}!</p>
{:catch error}
  <p>Cannot load data</p>
  <p>{error.message}</p>
{/await}
