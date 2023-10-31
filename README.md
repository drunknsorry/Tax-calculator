# Tax-calculator

This is a basic get api implementation of a tax calculator which when provided a year and salary query, responds back with tax brackets, total tax, effective tax rate and related information.

To use this project, please clone or checkout the repo and 

#### Api documentation can be found at: https://documenter.getpostman.com/view/30818865/2s9YXbARSe

Logs can be found within the /log folder.

### Changes and improvements needed
- In a typical production environment, this API would normally make requests to the test API in the background and store the necessary data in a temporary data store like Redis or Memcached. However, I haven't implemented this feature because based on the structure of the assessment; my assumption was the requests need to be made individually.
- The package uses go standard libraries as much as possible. This is due to go's promise of backwards compatibility; this ensures as the api is developed and expanded further, breaking changes do not occur between version updates. Alternative modules to use would be gorilla mux, chi or gin, logrus etc.
- The `api` package can be further subdivided to separate the route handling from the logic, but due to time constraints, I couldn't take this step.
- I introduced a `logger` package to reduce redundant code for logging. However, the current implementation needs improvement. While it works during regular usage, it fails in tests due to the way logs are created, opened, and written. Unfortunately, time constraints prevented me from refining this further.
- There's room for improvement in the logging and `logger` package itself. The current log mainly focuses on errors, but it should ideally include a detailed account of every action the server takes. Unfortunately, I couldn't enhance this aspect due to time limitations.
- My current implementation can be horizontally scaled using Docker containers or container clusters. I initially planned to provide a Docker image, but time constraints got in the way. If you intend to scale on bare metal (with more CPU cores), I'll need to implement goroutines to significantly speed up data processing.
- Testing also needs some work. The current version predates the `logger` package implementation, which leads to some unfortunate failures as mentioned above. I need to expand my testing coverage to include more scenarios and data to ensure a more robust and reliable system.
