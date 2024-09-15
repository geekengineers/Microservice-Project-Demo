import LogoImage from "../assets/images/logo.png"

export function HeaderComponent() {
    return (
        <div className="flex justify-center">
            <div dir="rtl" className="container bg-slate-500 navbar shadow-none bg-transparent">
                <div className="navbar-start">
                    <a className="navbar-item">
                        <img className="w-44" src={LogoImage} alt="Logo" />
                    </a>
                </div>
                <div className="navbar-center flex-1 flex flex-row bg-red-400">
                    <a className="navbar-item">موبایل</a>
                    <a className="navbar-item">لپتاپ و تبلت</a>
                    <a className="navbar-item">هدفون و هندزفری</a>
                    <a className="navbar-item">فروشگاه</a>
                    <a className="navbar-item">درباره ما</a>
                    <a className="navbar-item">تماس با ما</a>
                </div>
                <div className="navbar-end">
                    <a className="btn btn-primary">ورود</a>
                </div>
            </div>
        </div>
    )
}