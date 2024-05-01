<template>
  <main class="flex flex-col flex-grow m-2 border bg-gray-100">
    <h1 class="text-2xl font-bold text-center">Message</h1>
    <section class="m-6 h-[100%]">
      <!-- Usar la propiedad computada formattedBodyContent para mostrar el cuerpo del mensaje -->
      <p v-html="formattedBodyContent"></p>
    </section>
  </main>
</template>

  
  <script>
  export default {
    name: 'BodyComponent',
    props: ['bodyContent'],
    computed: {
      // Propiedad computada para formatear el cuerpo del mensaje con saltos de línea
      formattedBodyContent() {
        // Reemplazar los saltos de línea por etiquetas <br>
        return this.bodyContent.replace(/\n/g, '<br>');
      }
    }
  };
  </script>
  