import Logo from "../asserts/svg/logo.svg";

const Avatar = (props) => {
  const { src, className, ...rest } = props;

  return (
    <img
      src={src ? src : Logo}
      className={`w-full h-full ${className}`}
      {...rest}
    />
  );
};

export default Avatar;
