<template>
  <form class="w-full px-4" @submit.prevent="isSearching">
    <div>
      <input type="text" name="q"
        class="w-full border h-12 shadow p-4 rounded-full dark:text-gray-800 dark:border-gray-700 dark:bg-gray-200"
        placeholder="search" v-model="query" />
    </div>
  </form>
</template>

<script>
export default {
  name: 'SearchComponent',
  props: ['currentPage'],
  data() {
    return {
      query: '',
      booleanSearch: false,
      API_URL:process.env.VUE_APP_PATH_START,
    };
  },
  watch: {
    currentPage: {
      immediate: true,
      handler() {
        // Cuando currentPage cambia, llamar a handleSubmit
        this.handleSubmit();
      }
    }
  },
  methods: {
    async isSearching() {
      this.booleanSearch = true;
      this.$emit('searching', true) // Emitir el evento para actualizar currentPage en el componente padre
      this.handleSubmit();
    },

    async handleSubmit() {
      // Realizar la solicitud POST con los datos de búsqueda
      try {
        if (this.booleanSearch) {
          const response = await fetch(`${this.API_URL}search/${this.currentPage}`, {
            method: 'POST',
            headers: {
              'Content-Type': 'application/json',
            },
            body: JSON.stringify({
              term: this.query,
            }),
          });

          if (!response.ok) {
            throw new Error('Error en la solicitud');
          }

          const data = await response.json();
          this.$emit('search-results', data);
          this.$emit('page-change', data.total); // Emitir el evento page-change para actualizar la página actual en PageComponent
        }
        else {
          const response = await fetch(`${this.API_URL}searchall/${this.currentPage}`, {
            method: 'GET',
            headers: {
              'Content-Type': 'application/json',
            },
          });

          if (!response.ok) {
            throw new Error('Error en la solicitud');
          }

          const data = await response.json();
          this.$emit('search-results', data);
        }
        // Emitir los resultados de la búsqueda al componente padre
      } catch (error) {
        console.error(error);
      }
    },
  },
};
</script>

<style scoped>
/* Aquí puedes agregar tus estilos */
</style>
