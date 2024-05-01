<template>
  <div>
    <NavBar />
    <SearchComponent @search-results="handleSearchResults" />
    <div class="flex flex-row">
      <ResponsiveTable :items="tableData" @show-body="showBody" />
      <BodyComponent class="flex-grow max-w-[200px]:" :bodyContent="bodyContent" />
    </div>
  </div>
</template>

<script>
import NavBar from './components/Navbar.vue'
import ResponsiveTable from './components/ResponsiveTable.vue'
import SearchComponent from './components/SearchComponent.vue'
import BodyComponent from './components/BodyComponent.vue'

export default {
  name: 'App',
  components: {
    NavBar,
    SearchComponent,
    ResponsiveTable,
    BodyComponent,
  },
  data() {
    return {
      bodyContent: '', // Inicializamos bodyContent como una cadena vacía
      tableData: [], // Inicializamos tableData como un arreglo vacío para almacenar los datos de la tabla
    };
  },
  methods: {
    showBody(body) {
      this.bodyContent = body; // Actualizamos el valor de bodyContent con el cuerpo del mensaje
    },
    handleSearchResults(results) {
      this.tableData = results; // Actualizamos tableData con los resultados de la búsqueda
    }
  },
}
</script>
