<template>
  <div class="bg-gray-200">
    <NavBar />
    <SearchComponent :currentPage="currentPage" @searching="handleSearching" @search-results="handleSearchResults" />
    <PageComponent :pageSize="totalPage" :actualPage="currentPage" @page-reset="handleSearching"
      @page-change="currentPage = $event" />
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
import PageComponent from './components/PageComponent.vue'

export default {
  name: 'App',
  components: {
    NavBar,
    SearchComponent,
    TableComponent,
    BodyComponent,
    PageComponent

  },
  data() {
    return {
      bodyContent: '',
      tableData: [],
      totalPage: 0,
      currentPage: 1,
      API_URL:process.env.VUE_APP_PATH_START,
    };
  },
  methods: {

    handleSearching() {
      // Cuando se realiza una búsqueda, establecer currentPage en 1
      this.currentPage = 1;
    },

    showBody(body) {
      this.bodyContent = body;
    },
    handleSearchResults(results) {
      this.tableData = results;
      this.totalPage = results.total;
      if (results.total % 20){
        this.totalPage = Math.floor(results.total / 20) + 1;
      }
      else{
        this.totalPage = results.total / 20;
      }
    },
    async deleteItem(id) {
  try {
    const response = await fetch(`${this.API_URL}delete/${id}`, {
      method: 'DELETE'
    });
    if (!response.ok) {
      throw new Error('Error al eliminar el elemento');
    }
    // Agregar control de tipo para asegurarse de que tableData sea un array
    if (Array.isArray(this.tableData.emails)) {
      this.tableData.emails = this.tableData.emails.filter(item => item._id !== id);
      console.log('Email eliminado');
    } else {
      console.error('tableData is not an array');
    }
  } catch (error) {
    console.error(error);
  }
}

  },
}
</script>
