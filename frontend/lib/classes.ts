const filterClass = (name: string) => name?.length > 0;

type CssClass = string | number | undefined | boolean;

export const classes = (props: CssClass[]) => {
    return props.filter(filterClass as any).join(' ');
};

export default classes;
