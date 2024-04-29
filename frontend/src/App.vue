<template>
  <NavBar />
  <SearchComponent />
  <div class="grid grid-cols-2 gap-2 mt-2 w-screen overflow-y-scroll p-2">
    <ResponsiveTable />
    <main class="border border-gray-300 shadow-lg mt-2">
      <h2 class="text-2xl font-bold text-center">Main Content</h2>
      Lorem ipsum dolor sit amet consectetur adipisicing elit. Accusantium quam
      eos aperiam est inventore soluta dignissimos. Cumque ut rerum itaque, qui
      deleniti animi consequatur quia pariatur minima sed fugit quam!
    </main>
  </div>
</template>

<script>
import NavBar from './components/Navbar.vue';
import ResponsiveTable from './components/ResponsiveTable.vue';
import SearchComponent from './components/SearchComponent.vue';
export default {
  name: 'App',
  components: {
    NavBar,
    SearchComponent,
    ResponsiveTable,
  },
};
</script>
