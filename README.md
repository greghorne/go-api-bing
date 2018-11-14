# go-api-bing

Go: API that handles requests for isochrones from Bing Maps API.  The returned JSON is ready for use in LeafletJS.com

- This was written for using in a different project and the functionality is narrow in scope.
- The API returns JSON in text format of only the GeoJSON portion of what Bing returns.

__*Deployment:*__ *http://zotact1.ddns.net:8001/v1/bing-isochrone/{lng}/{lat}/{time}/{key}*

- __*lng*__ => longitude (decimal degrees)
- __*lat*__ => latitude (decimal degrees)
- __*time*__ => drive time polygon in seconds
- __*key*__ => Bing key