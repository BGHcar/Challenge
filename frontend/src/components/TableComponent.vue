<template>
  <div class="w-full overflow-auto min-w-[100%] mine:max-w-[50%] mine:min-w-[50%]">
    <table
      class="text-[12px] flex-col flex-grow min-w-[98%] divide-gray-200 border m-2 border-gray-300 shadow-lg bg-gray-100">
      <!-- Encabezado de la tabla -->
      <thead class="bg-gray-50 divide-y">
        <tr class="text-left">
          <th class="border px-3 py-3 font-medium text-gray-500 uppercase ">From</th>
          <th class="border px-3 py-3 font-medium text-gray-500 uppercase ">To</th>
          <th class="border px-3 py-3 font-medium text-gray-500 uppercase ">Subject</th>
          <th class="border px-3 py-3 font-medium text-gray-500 uppercase ">Actions</th>
        </tr>
      </thead>
      <!-- Cuerpo de la tabla -->
      <tbody class=" divide-y divide-gray-200 ">
        <tr class="hover:bg-indigo-200" v-for="(item, index) in items.emails" :key="index">
          <td class="border px-3 py-2">{{ item.From }}</td>
          <td class="border px-3 py-2">{{ item.To }}</td>
          <td class="border px-3 py-2">{{ item.Subject }}</td>
          <td class="border px-3 py-2 flex justify-center">
            <!-- Usar item.Message directamente en lugar de item.emails[Message] -->
            <button class="bg-indigo-500 focus:ring-4 hover:bg-indigo-700 text-white font-bold py-2 w-[100px] mr-2 rounded mb-2"
              @click="isDataShow(this.showDate, item) ">
              VIEW
            </button>
            <!-- Pasar solo el ID del correo electrónico a confirmDelete -->
            <button class="bg-gray-500 focus:ring-4 hover:bg-gray-700 text-white py-2 w-[100px] mb-2 font-bold rounded"
              @click="confirmDelete(item._id)">
              DELETE
            </button>
          </td>
        </tr>
      </tbody>
    </table>

    <!-- Modal para confirmación de eliminación -->
    <div v-if="showModal" class="fixed inset-0 z-10 overflow-y-auto" aria-labelledby="modal-title" role="dialog"
      aria-modal="true">
      <div class="flex items-end justify-center min-h-screen pt-4 px-4 pb-20 text-center sm:block sm:p-0">
        <div class="fixed inset-0 bg-gray-500 bg-opacity-75 transition-opacity" aria-hidden="true"></div>

        <!-- This element is to trick the browser into centering the modal contents. -->
        <span class="hidden sm:inline-block sm:align-middle sm:h-screen" aria-hidden="true">&#8203;</span>

        <div
          class="inline-block align-bottom bg-white rounded-lg text-left overflow-hidden shadow-xl transform transition-all sm:my-8 sm:align-middle sm:max-w-lg sm:w-full"
          role="dialog" aria-labelledby="modal-title" aria-describedby="modal-description">
          <div class="bg-white px-4 pt-5 pb-4 sm:p-6 sm:pb-4">
            <div class="sm:flex sm:items-start">
              <div
                class="mx-auto flex-shrink-0 flex items-center justify-center h-12 w-12 rounded-full bg-red-100 sm:mx-0 sm:h-10 sm:w-10">
                <!-- Heroicon name: outline/exclamation -->
                <svg class="h-6 w-6 text-red-600" xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24"
                  stroke="currentColor" aria-hidden="true">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2"
                    d="M12 9v2m0 4h.01M12 17h.01M3 4h18a2 2 0 012 2v12a2 2 0 01-2 2H3a2 2 0 01-2-2V6a2 2 0 012-2z" />
                </svg>
              </div>
              <div class="mt-3 text-center sm:mt-0 sm:ml-4 sm:text-left">
                <h3 class="text-lg leading-6 font-medium text-gray-900" id="modal-title">
                  Delete item
                </h3>
                <div class="mt-2">
                  <p class="text-sm text-gray-500" id="modal-description">
                    Are you sure you want to delete this item? This action cannot be undone.
                  </p>
                </div>
              </div>
            </div>
          </div>
          <div class="bg-gray-50 px-4 py-3 sm:px-6 sm:flex sm:flex-row-reverse">
            <button type="button" @click="deleteItem"
              class="w-full inline-flex justify-center rounded-md border border-transparent shadow-sm px-4 py-2 bg-red-600 text-base font-medium text-white hover:bg-red-700 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-red-500 sm:ml-3 sm:w-auto sm:text-sm">
              Delete
            </button>
            <button type="button" @click="closeModal"
              class="mt-3 w-full inline-flex justify-center rounded-md border border-gray-300 shadow-sm px-4 py-2 bg-white text-base font-medium text-gray-700 hover:bg-gray-50 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-indigo-500 sm:mt-0 sm:ml-3 sm:w-auto sm:text-sm">
              Cancel
            </button>
          </div>
        </div>
      </div>
    </div>

  </div>
</template>

<script>
export default {
  name: 'TableComponent',
  props: ['items'], // Recibe los datos de la tabla como prop 'items' desde el componente padre (PageComponent)
  data() {
    return {
      showModal: false, // Controla la visibilidad del modal
      itemIdToDelete: null, // Almacena el ID del elemento a eliminar
      showDate: false,  // Controla la visibilidad del cuerpo del mensaje
      messageId: '',  // Almacena el ID del mensaje actual
    };
  },
  methods: {

    isDataShow(flag, body) {
      if (!flag || this.messageId !== body._id) {
        this.messageId = body._id;
        this.showDate = true;
        this.showBodyHandler(body.Message);
      } else {
        this.showBodyHandler('');
        this.showDate = false;
      }

    },

    showBodyHandler(body) {
      this.$emit('show-body', body);
    },
    confirmDelete(id) {
      // Mostrar el modal de confirmación y almacenar el ID del elemento a eliminar
      this.itemIdToDelete = id;
      this.showModal = true;
    },
    deleteItem() {
      // Emitir evento para solicitar eliminación al componente padre
      this.$emit('delete-item', this.itemIdToDelete);
      // Ocultar el modal de confirmación
      this.showModal = false;
    },
    closeModal() {
      // Ocultar el modal de confirmación
      this.showModal = false;
    }
  },
};
</script>
