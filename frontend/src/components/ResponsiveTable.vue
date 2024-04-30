<template>
  <div>
    <table class="text-[12px] divide-gray-200 mt-2 border border-gray-300 shadow-lg h-full w-full">
      <!-- Encabezado de la tabla -->
      <thead class="bg-gray-50">
        <tr class="text-center">
          <th class="ml-10 px-3 py-3 text-left font-medium text-gray-500 uppercase tracking-wider">
            From
          </th>
          <th class="px-3 py-3 text-left font-medium text-gray-500 uppercase tracking-wider">
            To
          </th>
          <th class="px-3 py-3 text-left font-medium text-gray-500 uppercase tracking-wider">
            Subject
          </th>
          <th class="px-3 py-3 text-left font-medium text-gray-500 uppercase tracking-wider">
            Actions
          </th>
        </tr>
      </thead>
      <!-- Cuerpo de la tabla -->
      <tbody class="bg-white divide-y divide-gray-200">
        <tr class=" hover:bg-indigo-200" v-for="(item, index) in items" :key="index">
          <td class="px-6 py-4  whitespace-nowrap">{{ item.email }}</td>
          <td class="px-6 py-4 whitespace-nowrap">{{ item.email }}</td>
          <td class="px-6 py-4 whitespace-nowrap">{{ item.name }}</td>
          <td class="px-6 py-4 whitespace-nowrap">
            <button class="bg-blue-500 hover:bg-blue-700 text-black font-bold py-2 px-6 rounded mr-4">
              VIEW
            </button>
            <button class="bg-red-500 hover:bg-red-700 text-black font-bold py-2 px-6 rounded">
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
  data() {
    return {
      items: [],
    };
  },
  methods: {
    async getData() {
      try {
        const data = await fetch('https://jsonplaceholder.typicode.com/users');
        this.items = await data.json();
      } catch (error) {
        console.error(error);
      }
    },
  },
  created() {
    this.getData();
  },
};
</script>
