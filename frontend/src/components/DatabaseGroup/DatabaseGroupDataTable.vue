<template>
  <NDataTable
    v-bind="$attrs"
    size="small"
    :columns="columnList"
    :data="data"
    :striped="true"
    :bordered="bordered"
    :loading="loading"
    :row-key="(data: ComposedDatabaseGroup) => data.name"
    :checked-row-keys="selectedDatabaseGroupNames"
    :row-props="rowProps"
    :pagination="{ pageSize: 20 }"
    :paginate-single-page="false"
    @update:checked-row-keys="
      (val) => $emit('update:selected-database-group-names', val as string[])
    "
  />
</template>

<script lang="tsx" setup>
import { ExternalLinkIcon } from "lucide-vue-next";
import { NDataTable, NEllipsis, type DataTableColumn } from "naive-ui";
import { computed } from "vue";
import { useI18n } from "vue-i18n";
import { useRouter } from "vue-router";
import { PROJECT_V1_ROUTE_DATABASE_GROUP_DETAIL } from "@/router/dashboard/projectV1";
import { getProjectNameAndDatabaseGroupName } from "@/store";
import type { ComposedDatabaseGroup } from "@/types";

type DatabaseGroupDataTableColumn = DataTableColumn<ComposedDatabaseGroup> & {
  hide?: boolean;
};

const props = withDefaults(
  defineProps<{
    databaseGroupList: ComposedDatabaseGroup[];
    bordered?: boolean;
    loading?: boolean;
    customClick?: boolean;
    showSelection?: boolean;
    showExternalLink?: boolean;
    singleSelection?: boolean;
    selectedDatabaseGroupNames?: string[];
  }>(),
  {
    bordered: true,
    selectedDatabaseGroupNames: () => [],
  }
);

const emit = defineEmits<{
  (
    event: "row-click",
    e: MouseEvent,
    databaseGroup: ComposedDatabaseGroup
  ): void;
  (event: "update:selected-database-group-names", val: string[]): void;
}>();

const { t } = useI18n();
const router = useRouter();

const columnList = computed((): DatabaseGroupDataTableColumn[] => {
  const rawColumnList: DatabaseGroupDataTableColumn[] = [
    {
      type: "selection",
      multiple: !props.singleSelection,
      hide: !props.showSelection,
      cellProps: () => {
        return {
          onClick: (e: MouseEvent) => {
            e.stopPropagation();
          },
        };
      },
    },
    {
      key: "title",
      title: t("common.name"),
      width: 256,
      ellipsis: true,
      resizable: true,
      render: (data) => {
        return (
          <div class="space-x-2">
            <span>{data.title}</span>
          </div>
        );
      },
    },
    {
      key: "expression",
      title: t("database.expression"),
      ellipsis: true,
      render: (data) => {
        if (!data.databaseExpr || data.databaseExpr.expression === "") {
          return <span class="textinfolabel italic">{t("common.empty")}</span>;
        }
        return <NEllipsis>{data.databaseExpr.expression}</NEllipsis>;
      },
    },
    {
      key: "externalLink",
      title: "",
      width: 48,
      hide: !props.showExternalLink,
      render: (data) => {
        const openExternalLink = (e: MouseEvent) => {
          e.stopPropagation();
          const [projectId, databaseGroupName] =
            getProjectNameAndDatabaseGroupName(data.name);
          const url = router.resolve({
            name: PROJECT_V1_ROUTE_DATABASE_GROUP_DETAIL,
            params: {
              projectId,
              databaseGroupName,
            },
          }).fullPath;
          window.open(url, "_blank");
        };

        return (
          <div
            class="flex items-center justify-end cursor-pointer w-6 h-6 p-1 opacity-60 hover:opacity-100 hover:bg-white hover:shadow rounded"
            onClick={openExternalLink}
          >
            <ExternalLinkIcon class="w-4 h-auto" />
          </div>
        );
      },
    },
  ];

  return rawColumnList.filter((column) => !column.hide);
});

const data = computed(() => {
  return [...props.databaseGroupList];
});

const rowProps = (databaseGroup: ComposedDatabaseGroup) => {
  return {
    style: "cursor: pointer;",
    onClick: (e: MouseEvent) => {
      if (props.customClick) {
        emit("row-click", e, databaseGroup);
        return;
      }

      if (props.singleSelection) {
        emit("update:selected-database-group-names", [databaseGroup.name]);
      } else {
        const selectedDatabaseGroupNameList = new Set(
          props.selectedDatabaseGroupNames
        );
        if (selectedDatabaseGroupNameList.has(databaseGroup.name)) {
          selectedDatabaseGroupNameList.delete(databaseGroup.name);
        } else {
          selectedDatabaseGroupNameList.add(databaseGroup.name);
        }
        emit("update:selected-database-group-names", [
          ...selectedDatabaseGroupNameList,
        ]);
      }
    },
  };
};
</script>
