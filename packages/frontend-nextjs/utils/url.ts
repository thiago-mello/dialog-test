/**
 * Converts an object of parameters into a URL query string
 * @param params - Object containing key-value pairs to convert to query string
 * @returns URL encoded query string
 * @example
 * // Returns "foo=bar&baz=123"
 * getUrlQueryStringFromParams({foo: "bar", baz: 123})
 */
export function getUrlQueryStringFromParams(params: Object) {
  const seachParams = new URLSearchParams();
  Object.entries(params).forEach(([key, value]) => {
    if (value || value === false) {
      seachParams.append(key, value.toString());
    }
  });

  return seachParams.toString();
}
