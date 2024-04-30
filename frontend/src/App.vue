<template>
  <NavBar />
  <SearchComponent />
  <div class="flex flex-col gap-2 mt-2 p-2 md:flex md:flex-row bg-zinc-600 h-full">
    <ResponsiveTable/>
    <BodyComponent/>

  </div>
</template>

<script>
import NavBar from './components/Navbar.vue';
import ResponsiveTable from './components/ResponsiveTable.vue';
import SearchComponent from './components/SearchComponent.vue';
import BodyComponent from './components/BodyComponent.vue';
export default {
  name: 'App',
  components: {
    NavBar,
    SearchComponent,
    ResponsiveTable,
    BodyComponent,
  },
};
</script>
