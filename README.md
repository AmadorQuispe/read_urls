# URL Filter and Counter in Go

Este programa lee un archivo de texto plano con una lista de URLs, filtra aquellas que contienen la palabra "shop" en su dominio y que terminan con la extensión ".html", eliminando duplicados y contando cuántas cumplen con estos criterios.

## Requisitos

Para ejecutar este programa, necesitarás tener instalado:

- **Go (1.18 o superior)**

## Instrucciones de Ejecución

1. **Clonar el repositorio o descargar los archivos:**
   
   Descarga el archivo fuente `main.go` y el archivo `urls.txt` que contiene las URLs para ser procesadas.

2. **Colocar las URLs:**

   Asegúrate de que las URLs estén almacenadas en un archivo llamado `urls.txt`. Aquí tienes un ejemplo del formato:

   ```txt
   https://myshop.com/products/item123.html
   https://shoponline.com/deals/todaysdeal.html
   http://nonshop.com/specials.html
   https://anotherdomain.com/about.html
   https://shopify.com/store/productX.html
   ```

3. **Compilar el programa:**

   Para compilar el programa, abre una terminal en el directorio donde se encuentra el archivo `main.go` y ejecuta:

   ```bash
   go build -o url_filter
   ```

4. **Ejecutar el programa:**

   Una vez que se haya compilado el programa, puedes ejecutarlo con el siguiente comando:

   ```bash
   ./url_filter
   ```

5. **Ver los resultados:**

   El programa mostrará el número total de URLs que cumplen con los criterios de filtrado (contienen "shop" en el dominio y terminan en ".html") y imprime todas las URLs filtradas.

## Algoritmo

El algoritmo está diseñado para ser eficiente en tiempo y uso de memoria, incluso cuando se trabaja con millones de URLs:

1. **Lectura de archivo eficiente**: Utiliza `bufio.Scanner` para leer el archivo línea por línea, lo que permite manejar archivos de gran tamaño sin cargar todo en memoria.
   
2. **Verificación de criterios**: En lugar de usar expresiones regulares costosas, se usa `strings.Contains()` para verificar si el dominio contiene "shop" y `strings.HasSuffix()` para verificar si la URL termina en ".html".

3. **Eliminación de duplicados**: Se utiliza un mapa (`map[string]struct{}`) para almacenar solo las URLs únicas y así eliminar duplicados de manera eficiente.

4. **Escalabilidad**: La estructura y las técnicas utilizadas permiten que el programa maneje archivos con millones de URLs sin problemas de rendimiento ni de memoria.

## Optimización

- **Uso de `map[string]struct{}`**: En lugar de usar `map[string]bool`, se utiliza una estructura vacía para minimizar el uso de memoria al eliminar duplicados.
- **Procesamiento en línea**: El programa procesa una URL a la vez, sin cargar todo el archivo en memoria, lo que permite trabajar con archivos extremadamente grandes.

