<template>
  <div class="w-full overflow-auto min-w-[100%] mine:max-w-[50%] mine:min-w-[50%]">
    <table class="max-w-[50%] sm:max-w-full text-[12px] flex flex-col flex-grow divide-gray-200 border m-2 border-gray-300 shadow-lg">
      <!-- Encabezado de la tabla -->
      <thead class="bg-gray-50">
        <tr class="text-left">
          <th class="border px-3 py-3 font-medium text-gray-500 uppercase ">From</th>
          <th class="border px-3 py-3 font-medium text-gray-500 uppercase ">To</th>
          <th class="border px-3 py-3 font-medium text-gray-500 uppercase ">Subject</th>
          <th class="border px-3 py-3 font-medium text-gray-500 uppercase ">Actions</th>
        </tr>
      </thead>
      <!-- Cuerpo de la tabla -->
      <tbody class="bg-white divide-y divide-gray-200">
        <tr class="hover:bg-indigo-300" v-for="(item, index) in items" :key="index">
          <td class="border px-3 py-2">{{ item.From }}</td>
          <td class="border px-3 py-2">{{ item.To }}</td>
          <td class="border px-3 py-2">{{ item.Subject }}</td>
          <td class="border px-3 py-2">
            <button class="bg-indigo-500 hover:bg-indigo-700 text-black font-bold py-2 w-[80px] mr-2 rounded"
              @click="showBodyHandler(item.Message, console.log(item.Message))">
              VIEW
            </button>
            <button class="bg-red-500 hover:bg-red-700 text-black py-2 w-[80px] font-bold rounded">
              DELETE
            </button>
          </td>
        </tr>
      </tbody>
    </table>

  </div>
</template>

<script>
export default {
  name: 'ResponsiveTable',
  props: ['items'], // Recibe los datos de la tabla como prop
  methods: {
    showBodyHandler(body) {
      this.$emit('show-body', body);
    }
  },
};
</script>
