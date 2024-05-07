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

      </div>
    </div>
  </nav>
</template>

<script>
export default {
  name: 'NavBar',
};
</script>
