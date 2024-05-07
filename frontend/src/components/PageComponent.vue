<template>
  <div>
    <!-- Botones de paginación -->
    <div class="flex justify-center items-center mt-4 mb-4">
      <button type="button" class="bg-indigo-200 rounded-lg hover:bg-indigo-400 mx-4 p-4" @click="previousPage">Anterior</button>
      <span v-for="pageNumber in visiblePageNumbers" :key="pageNumber">
        <button type="button" class="mx-2 p-2"
          :class="{ 'font-bold text-indigo-800': parseInt(pageNumber) === parseInt(this.currentPage) }"
          @click="goToPage(pageNumber)">{{ pageNumber }}</button>
      </span>
      <button type="button" class="bg-indigo-200 rounded-lg hover:bg-indigo-400 mx-4 p-4" @click="nextPage">Siguiente</button>
    </div>
  </div>
</template>

<script>


export default {
  name: 'PageComponent',
  props: ['pageSize', 'actualPage'],

  data() {
    return {
      currentPage: parseInt(this.actualPage),
    };
  },
  computed: {
    visiblePageNumbers() {
      const totalPages = this.pageSize;  // Obtén el número total de páginas
      return this.generatePageNumbers(totalPages);
    },
  },
  methods: {
    generatePageNumbers(totalPages) {
      if (totalPages <= 7) {
        return Array.from({ length: totalPages }, (_, i) => i + 1);  // Genera un array de 1 a totalPages
      } else if (this.currentPage <= 4) {
        return [1, 2, 3, 4, 5, 6, 7];  // Si currentPage es menor o igual a 4, muestra los primeros 7 números
      } else if (this.currentPage >= totalPages - 3) {
        return [totalPages - 6, totalPages - 5, totalPages - 4, totalPages - 3, totalPages - 2, totalPages - 1, totalPages]; // Si currentPage es mayor o igual a totalPages - 3, muestra los últimos 7 números
      } else { 
        return [this.currentPage - 3, this.currentPage - 2, this.currentPage - 1, this.currentPage, this.currentPage + 1, this.currentPage + 2, this.currentPage + 3]; // En cualquier otro caso, muestra los números alrededor de currentPage
      }
    },
    async nextPage() {
      if (this.currentPage < this.pageSize){
      this.currentPage++;
      this.$emit('page-change', this.currentPage);  // Emitir el evento para actualizar la página actual en el componente padre
      }
    },
    async previousPage() {
      if (this.currentPage > 1) {
        this.currentPage--;
        this.$emit('page-change', this.currentPage); 
      }
    },
    async goToPage(pageNumber) {
      this.currentPage = pageNumber;
      this.$emit('page-change', this.currentPage);  
    },
    async resetPage() {
      this.currentPage = parseInt(this.actualPage); // convertir actualPage a entero
      this.$emit('page-change', this.currentPage);
    }
  },
  watch: {
    // Observa cambios en actualPage y actualiza currentPage
    actualPage: {
      immediate: true, // Para que se ejecute la primera vez
      handler(newVal) {
        this.currentPage = parseInt(newVal); // convertir newVal a entero
      }
    }
  }
};
</script>
