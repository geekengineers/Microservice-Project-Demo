import { Swiper, SwiperSlide } from "swiper/react"

export function HomeComponent() {
    return (
        <>
            <Swiper
                id="home-slider"
                slidesPerView={1}
                spaceBetween={10}
            >
                <SwiperSlide>
                    <div className="swiper-slider-metadata">
                        <h1 className="text-5xl font-bold">Product Name</h1>
                        <p className="relative left-10 mt-4 w-96">Magna laborum labore et laboris Lorem amet. Laborum sunt sit ad enim aliquip Lorem enim aute enim pariatur ex in consequat. Tempor dolore dolor amet pariatur irure dolore non non exercitation consequat cillum et.</p>

                    </div>
                    <img src="https://png.pngtree.com/thumb_back/fh260/background/20201101/pngtree-mock-up-podium-for-product-presentation-blue-background-3d-render-illustration-image_452966.jpg" />
                </SwiperSlide>
                <SwiperSlide>
                    <img src="https://png.pngtree.com/thumb_back/fh260/background/20231002/pngtree-realistic-product-presentation-stunning-blue-themed-3d-podium-design-for-image_13556182.png" />
                </SwiperSlide>
                <SwiperSlide>
                    <img src="https://png.pngtree.com/thumb_back/fh260/background/20230707/pngtree-d-render-of-a-glass-podium-stand-on-a-futuristic-blue-image_3814625.jpg" />
                </SwiperSlide>
            </Swiper>
        </>
    )
}