<template>
  <table
    class="min-w-full divide-y divide-gray-200 mt-2 border border-gray-300 shadow-lg"
  >
    <thead class="bg-gray-50">
      <tr class="text-center">
        <th
          class="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider"
        >
          From
        </th>
        <th
          class="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider"
        >
          To
        </th>
        <th
          class="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider"
        >
          Subject
        </th>
      </tr>
    </thead>
    <tbody class="bg-white divide-y divide-gray-200">
      <tr v-for="(item, index) in items" :key="index">
        <td class="px-6 py-4 whitespace-nowrap">{{ item.email }}</td>
        <td class="px-6 py-4 whitespace-nowrap">{{ item.email }}</td>
        <td class="px-6 py-4 whitespace-nowrap">{{ item.name }}</td>
      </tr>
    </tbody>
  </table>
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
