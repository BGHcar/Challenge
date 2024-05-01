<template>
  <div class="bg-gray-200">
    <NavBar />
    <SearchComponent @search-results="handleSearchResults" />
    <div class="flex flex-col mine:flex-row">
      <TableComponent :items="tableData" @show-body="showBody" @delete-item="deleteItem" />
      <BodyComponent :bodyContent="bodyContent" />
    </div>
  </div>
</template>

<script>
import NavBar from './components/Navbar.vue'
import TableComponent from './components/TableComponent.vue'
import SearchComponent from './components/SearchComponent.vue'
import BodyComponent from './components/BodyComponent.vue'

export default {
  name: 'App',
  components: {
    NavBar,
    SearchComponent,
    TableComponent,
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
    },
    async deleteItem(id) {
      try {
        const response = await fetch(`http://localhost:9000/delete/${id}`, {
          method: 'DELETE'
        });
        if (!response.ok) {
          throw new Error('Error al eliminar el elemento');
        }
        // Actualizar la lista de elementos después de eliminar
        this.tableData = this.tableData.filter(item => item._id !== id);
      } catch (error) {
        console.error(error);
      }
    }
  },
}
</script>
