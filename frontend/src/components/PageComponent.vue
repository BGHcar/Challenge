<template>
  <div>
    <!-- Botones de paginación -->
    <div class="flex justify-center items-center mt-4 mb-4">
      <button class="bg-indigo-200 rounded-lg hover:bg-indigo-400 mx-4 p-4" @click="previousPage"
        :disabled="currentPage === 1">Anterior</button>
      <span v-for="pageNumber in visiblePageNumbers" :key="pageNumber">
        <button class="mx-2 p-2"
          :class="{ 'font-bold text-indigo-800': parseInt(pageNumber) === parseInt(this.currentPage) }"
          @click="goToPage(pageNumber)">{{ pageNumber }}</button>
      </span>
      <button class="bg-indigo-200 rounded-lg hover:bg-indigo-400 mx-4 p-4" @click="nextPage">Siguiente</button>
    </div>
  </div>
</template>

<script>

export default {
  name: 'PageComponent',
  props: ['totalEmails', 'pageSize', 'actualPage'],

  data() {
    return {
      currentPage: parseInt(this.actualPage)
    };
  },
  computed: {
    visiblePageNumbers() {
      const totalPages = Math.ceil(this.totalEmails / this.pageSize);
      return this.generatePageNumbers(totalPages);
    },
  },
  methods: {
    generatePageNumbers(totalPages) {
      if (totalPages <= 7) {
        return Array.from({ length: totalPages }, (_, i) => i + 1);
      } else if (this.currentPage <= 4) {
        return [1, 2, 3, 4, 5, 6, 7];
      } else if (this.currentPage >= totalPages - 3) {
        return [totalPages - 6, totalPages - 5, totalPages - 4, totalPages - 3, totalPages - 2, totalPages - 1, totalPages];
      } else {
        return [this.currentPage - 3, this.currentPage - 2, this.currentPage - 1, this.currentPage, this.currentPage + 1, this.currentPage + 2, this.currentPage + 3];
      }
    },
    async nextPage() {
      this.currentPage++;
      this.$emit('page-change', this.currentPage);
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
      this.currentPage = parseInt(this.actualPage); // Asegúrate de convertir actualPage a entero
      this.$emit('page-change', this.currentPage);
    }
  },
  watch: {
    // Observa cambios en actualPage y actualiza currentPage
    actualPage: {
      immediate: true, // Para que se ejecute la primera vez
      handler(newVal) {
        this.currentPage = parseInt(newVal); // Asegúrate de convertir newVal a entero
      }
    }
  }
};
</script>
