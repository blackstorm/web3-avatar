const imageReader = (file, onLoad) => {
  const reader = new FileReader();
  reader.addEventListener(
    "load",
    () => {
      onLoad(reader.result);
    },
    false
  );
  reader.readAsDataURL(file);
};

export { imageReader };
