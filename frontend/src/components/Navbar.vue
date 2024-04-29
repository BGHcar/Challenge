<!--./src/components/NavBar.vue -->

<template>
  <nav
    class="relative flex flex-wrap items-center justify-between px-2 py-3 bg-indigo-500 mb-3"
  >
    <div
      class="container px-4 mx-auto flex flex-wrap items-center justify-between"
    >
      <div
        class="w-full relative flex justify-between lg:w-auto px-4 lg:static lg:block lg:justify-start"
      >
        <a
          class="text-sm font-bold leading-relaxed inline-block mr-4 py-2 whitespace-nowrap uppercase text-white"
          href="#pablo"
        >
          <h1 class="text-sm font-bold text-gray-800 text-xl">
            Carta Searcher
          </h1>
        </a>
        <button
          class="text-white cursor-pointer text-xl leading-none px-3 py-1 border border-solid border-transparent rounded bg-transparent block lg:hidden outline-none focus:outline-none"
          type="button"
          v-on:click="toggleNavbar()"
        >
          <i class="fas fa-bars"></i>
        </button>
      </div>
      <div
        v-bind:class="{ hidden: !showMenu, flex: showMenu }"
        class="lg:flex lg:flex-grow items-center"
      ></div>
    </div>
  </nav>
</template>

<script>
export default {
  name: 'NavBar',
};
</script>
