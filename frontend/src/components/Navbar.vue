<!--./src/components/NavBar.vue -->

<template>
  <nav
    class="flex px-2 py-3 bg-indigo-500 mb-3 min-w-[100%]"
  >
    <div
      class="container px-4 flex flex-wrap items-center justify-between"
    >
      <div
        class="w-full relative flex lg:w-auto lg:static lg:block lg:justify-start"
      >
        <a
          class="text-sm font-bold leading-relaxed inline-block mr-4 whitespace-nowrap uppercase text-white"
        >
          <h1 class="text-white text-xl">
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
