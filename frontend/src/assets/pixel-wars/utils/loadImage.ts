export function loadImage(src: string): Promise<HTMLImageElement> {
  return new Promise((resolve, reject) => {
    const image = new Image();
    image.src = src;
    image.onload =  async () => resolve(image);
    image.onerror = async (err) => reject(err);
  });
}