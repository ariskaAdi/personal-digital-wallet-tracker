import React from "react";

type CardComponentProps = React.ComponentProps<"div"> & {
  title: string;
  titleClassName?: string;
  actions?: React.ReactNode;
};

const CardComponent = ({
  className,
  children,
  title,
  titleClassName,
  actions,
}: CardComponentProps) => {
  // Default classnames
  const defaultCardClassname =
    "bg-white dark:bg-slate-800 rounded-2xl p-6 shadow-sm border border-slate-200 dark:border-slate-700";
  const defaultTitleClassname =
    "text-sm font-semibold text-slate-600 dark:text-slate-400";

  const showHeader = title || actions;

  return (
    <div className={className ?? defaultCardClassname}>
      {showHeader && (
        <div className="flex items-center justify-between mb-6">
          <h2 className={titleClassName ?? defaultTitleClassname}>{title}</h2>
          {actions && <div className="flex items-center gap-2">{actions}</div>}
        </div>
      )}

      {children}
    </div>
  );
};

export default CardComponent;
